import { RefObject } from "react";
import { remoteVideoRef, textForSendSdpRef, textToReceiveSdpRef } from "./App";

let localStream: MediaStream;
let peerConnection: RTCPeerConnection | null;
let negotiationneededCounter = 0;

// シグナリングサーバへ接続する
const wsUrl = "ws://localhost:3001/";
const ws = new WebSocket(wsUrl);
ws.onopen = (evt) => {
  console.log("ws open()");
  console.log(evt);
};
ws.onerror = (err) => {
  console.error("ws onerror() ERR:", err);
};
ws.onmessage = async (evt) => {
  console.log("ws onmessage() data:", evt.data);
  let messageText: string;

  if (evt.data instanceof Blob) {
    try {
      messageText = await evt.data.text();
    } catch (error) {
      console.error("Error reading Blob data:", error);
      return;
    }
  } else if (typeof evt.data === "string") {
    messageText = evt.data;
  } else {
    console.error("Unsupported message data type:", typeof evt.data);
    return;
  }
  const message = JSON.parse(messageText);
  switch (message.type) {
    case "offer": {
      console.log("Received offer ...");
      const textToReceiveSdp = textToReceiveSdpRef.current;
      if (textToReceiveSdp !== null) {
        textToReceiveSdp.value = message.sdp;
      }
      setOffer(message);
      break;
    }
    case "answer": {
      console.log("Received answer ...");
      const textToReceiveSdp = textToReceiveSdpRef.current;
      if (textToReceiveSdp) {
        textToReceiveSdp.value = message.sdp;
      }
      setAnswer(message);
      break;
    }
    case "candidate": {
      console.log("Received ICE candidate ...");
      const candidate = new RTCIceCandidate(message.ice);
      console.log(candidate);
      addIceCandidate(candidate);
      break;
    }
    case "close": {
      console.log("peer is closed ...");
      const textForSendSdp = textForSendSdpRef.current;
      const textToReceiveSdp = textToReceiveSdpRef.current;
      hangUp(textForSendSdp, textToReceiveSdp);
      break;
    }
    default: {
      console.log("Invalid message");
      break;
    }
  }
};

// ICE candaidate受信時にセットする
function addIceCandidate(candidate: RTCIceCandidate) {
  if (peerConnection) {
    peerConnection.addIceCandidate(candidate);
  } else {
    console.error("PeerConnection not exist!");
    return;
  }
}

// ICE candidate生成時に送信する
function sendIceCandidate(candidate: RTCIceCandidate) {
  console.log("---sending ICE candidate ---");
  const message = JSON.stringify({ type: "candidate", ice: candidate });
  console.log("sending candidate=" + message);
  ws.send(message);
}

// getUserMediaでカメラ、マイクにアクセス
export async function startVideo(localVideo: RefObject<HTMLVideoElement>) {
  try {
    localStream = await navigator.mediaDevices.getUserMedia({
      video: true,
      audio: false,
    });
    if (localVideo.current) {
      playVideo(localVideo.current, localStream);
    }
  } catch (err) {
    console.error("mediaDevice.getUserMedia() error:", err);
  }
}

// Videoの再生を開始する
export async function playVideo(
  element: HTMLVideoElement,
  stream: MediaStream
) {
  element.srcObject = stream;
  try {
    element.muted = true;
    await element.play();
  } catch (error) {
    console.log("error auto play:" + error);
  }
}

// WebRTCを利用する準備をする
function prepareNewConnection(isOffer: boolean) {
  const pc_config = {
    iceServers: [{ urls: "stun:stun.webrtc.ecl.ntt.com:3478" }],
  };
  const peer = new RTCPeerConnection(pc_config);

  // リモートのMediStreamTrackを受信した時
  const remoteVideo = remoteVideoRef.current;
  if (remoteVideo) {
    peer.ontrack = (evt) => {
      console.log("-- peer.ontrack()");
      playVideo(remoteVideo, evt.streams[0]);
    };
  }

  // ICE Candidateを収集したときのイベント
  peer.onicecandidate = (evt) => {
    if (evt.candidate) {
      console.log(evt.candidate);
      sendIceCandidate(evt.candidate);
    } else {
      console.log("empty ice event");
      // sendSdp(peer.localDescription);
    }
  };

  // Offer側でネゴシエーションが必要になったときの処理
  peer.onnegotiationneeded = async () => {
    try {
      if (isOffer) {
        if (negotiationneededCounter === 0) {
          const offer = await peer.createOffer();
          console.log("createOffer() succsess in promise");
          await peer.setLocalDescription(offer);
          console.log("setLocalDescription() succsess in promise");
          sendSdp(peer.localDescription);
          negotiationneededCounter++;
        }
      }
    } catch (err) {
      console.error("setLocalDescription(offer) ERROR: ", err);
    }
  };

  // ICEのステータスが変更になったときの処理
  peer.oniceconnectionstatechange = function () {
    console.log(
      "ICE connection Status has changed to " + peer.iceConnectionState
    );
    switch (peer.iceConnectionState) {
      case "closed":
      case "failed":
        if (peerConnection) {
          const textForSendSdp = textForSendSdpRef.current;
          const textToReceiveSdp = textToReceiveSdpRef.current;
          hangUp(textForSendSdp, textToReceiveSdp);
        }
        break;
      case "disconnected":
        break;
    }
  };

  // ローカルのMediaStreamを利用できるようにする
  if (localStream) {
    console.log("Adding local stream...");
    localStream
      .getTracks()
      .forEach((track) => peer.addTrack(track, localStream));
  } else {
    console.warn("no local stream, but continue.");
  }

  return peer;
}

// 手動シグナリングのための処理を追加する
function sendSdp(sessionDescription: RTCSessionDescription | null) {
  console.log("---sending sdp ---");
  const textForSendSdp = textForSendSdpRef.current;
  if (textForSendSdp && sessionDescription) {
    textForSendSdp.value = sessionDescription.sdp;
  }
  /*---
   textForSendSdp.focus();
   textForSendSdp.select();
   ----*/
  const message = JSON.stringify(sessionDescription);
  console.log("sending SDP=" + message);
  ws.send(message);
}

// Connectボタンが押されたらWebRTCのOffer処理を開始
export function connect() {
  if (!peerConnection) {
    console.log("make Offer");
    peerConnection = prepareNewConnection(true);
  } else {
    console.warn("peer already exist.");
  }
}

// Answer SDPを生成する
async function makeAnswer() {
  console.log("sending Answer. Creating remote session description...");
  if (!peerConnection) {
    console.error("peerConnection NOT exist!");
    return;
  }
  try {
    const answer = await peerConnection.createAnswer();
    console.log("createAnswer() succsess in promise");
    await peerConnection.setLocalDescription(answer);
    console.log("setLocalDescription() succsess in promise");
    sendSdp(peerConnection.localDescription);
  } catch (err) {
    console.error(err);
  }
}

// Receive remote SDPボタンが押されたらOffer側とAnswer側で処理を分岐
export function onSdpText(textToReceiveSdp: HTMLTextAreaElement | null) {
  if (textToReceiveSdp === null) {
    console.error("textToReceiveSdp NOT exist!");
    return;
  }
  const text = textToReceiveSdp.value;
  if (peerConnection) {
    console.log("Received answer text...");
    const answer = new RTCSessionDescription({
      type: "answer",
      sdp: text,
    });
    setAnswer(answer);
  } else {
    console.log("Received offer text...");
    const offer = new RTCSessionDescription({
      type: "offer",
      sdp: text,
    });
    setOffer(offer);
  }
  textToReceiveSdp.value = "";
}

// Offer側のSDPをセットする処理
async function setOffer(sessionDescription: RTCSessionDescription) {
  if (peerConnection) {
    console.error("peerConnection alreay exist!");
  }
  peerConnection = prepareNewConnection(false);
  try {
    await peerConnection.setRemoteDescription(sessionDescription);
    console.log("setRemoteDescription(answer) succsess in promise");
    makeAnswer();
  } catch (err) {
    console.error("setRemoteDescription(offer) ERROR: ", err);
  }
}

// Answer側のSDPをセットする場合
async function setAnswer(sessionDescription: RTCSessionDescription) {
  if (!peerConnection) {
    console.error("peerConnection NOT exist!");
    return;
  }
  try {
    await peerConnection.setRemoteDescription(sessionDescription);
    console.log("setRemoteDescription(answer) succsess in promise");
  } catch (err) {
    console.error("setRemoteDescription(answer) ERROR: ", err);
  }
}

// P2P通信を切断する
export function hangUp(
  textForSendSdp: HTMLTextAreaElement | null,
  textToReceiveSdp: HTMLTextAreaElement | null
) {
  if (peerConnection) {
    if (peerConnection.iceConnectionState !== "closed") {
      peerConnection.close();
      peerConnection = null;
      negotiationneededCounter = 0;
      const message = JSON.stringify({ type: "close" });
      console.log("sending close message");
      ws.send(message);
      const remoteVideo = remoteVideoRef.current;
      cleanupVideoElement(remoteVideo);

      if (textForSendSdp) {
        textForSendSdp.value = "";
      }
      if (textToReceiveSdp) {
        textToReceiveSdp.value = "";
      }
      return;
    }
  }
  console.log("peerConnection is closed.");
}

// ビデオエレメントを初期化する
function cleanupVideoElement(element: HTMLVideoElement | null) {
  if (element) {
    element.pause();
    element.srcObject = null;
  }
}
