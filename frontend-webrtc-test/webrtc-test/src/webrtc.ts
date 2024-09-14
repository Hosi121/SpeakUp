import { RefObject } from "react";
import { remoteVideoRef, textForSendSdpRef, textToReceiveSdpRef } from "./App";

let localStream: MediaStream;
let peerConnection: RTCPeerConnection | null;
let negotiationneededCounter = 0;

// シグナリングサーバへ接続する
//const wsUrl = "ws://localhost:3001/";
const wsUrl = "ws://10.70.174.103:3001/";
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
  if (peerConnection && peerConnection.remoteDescription) {
    peerConnection
      .addIceCandidate(candidate)
      .then(() => console.log("Added ICE candidate successfully"))
      .catch((err) =>
        console.error("Error adding received ICE candidate", err)
      );
  } else {
    console.warn(
      "Cannot add ICE candidate: peerConnection is not set up or remoteDescription is not set"
    );
  }
}

// ICE candidate生成時に送信する
function sendIceCandidate(candidate: RTCIceCandidate) {
  // 空のcandidateは送信しない
  if (!candidate || !candidate.candidate) {
    console.log("Skipping empty ICE candidate");
    return;
  }
  console.log("Sending ICE candidate", candidate);
  const message = JSON.stringify({ type: "candidate", ice: candidate });
  ws.send(message);
  console.log("ICE candidate sent:", message);
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

export async function startLocalStream(
  localVideo: RefObject<HTMLVideoElement>
) {
  try {
    localStream = await navigator.mediaDevices.getUserMedia({
      video: true,
      audio: true,
    });
    if (localVideo.current) {
      playVideo(localVideo.current, localStream);
    }
  } catch (err) {
    console.error("mediaDevice.getUserMedia() error:", err);
    throw err;
  }
}

// Videoの再生を開始する
export async function playVideo(
  element: HTMLVideoElement,
  stream: MediaStream
) {
  element.srcObject = stream;
  try {
    await element.play();
  } catch (error) {
    console.log("error auto play:" + error);
  }
}

// WebRTCを利用する準備をする
function prepareNewConnection(
  isOffer: boolean,
  remoteVideoRef: RefObject<HTMLVideoElement>
) {
  const pc_config = {
    iceServers: [
      { urls: "stun:stun.l.google.com:19302" },
      { urls: "stun:stun1.l.google.com:19302" },
      { urls: "stun:stun2.l.google.com:19302" },
      { urls: "stun:stun3.l.google.com:19302" },
      { urls: "stun:stun4.l.google.com:19302" },
    ],
  };
  const peer = new RTCPeerConnection(pc_config);

  // リモートのMediaStreamTrackを受信した時
  peer.ontrack = (evt) => {
    console.log("-- peer.ontrack()", evt.track.kind);
    if (evt.track.kind === "video") {
      const remoteVideo = remoteVideoRef.current;
      if (remoteVideo) {
        if (remoteVideo.srcObject !== evt.streams[0]) {
          remoteVideo.srcObject = evt.streams[0];
          console.log("Received remote video stream");
        }
      }
    } else if (evt.track.kind === "audio") {
      // オーディオトラックの処理
      const remoteAudio = new Audio();
      remoteAudio.srcObject = evt.streams[0];
      remoteAudio.play();
      console.log("Remote audio playback started");
    }
  };

  // ICE Candidateを収集したときのイベント
  peer.onicecandidate = (evt) => {
    if (evt.candidate) {
      console.log("New ICE candidate: ", evt.candidate);
      sendIceCandidate(evt.candidate);
    } else {
      console.log("ICE gathering completed");
    }
  };

  // ICE接続状態が変化したときのイベント
  peer.oniceconnectionstatechange = () => {
    console.log("ICE connection state changed to: ", peer.iceConnectionState);
    switch (peer.iceConnectionState) {
      case "connected":
        console.log("ICE connected successfully");
        break;
      case "failed":
        console.log("ICE connection failed");
        // ここで接続の再試行ロジックを実装できます
        break;
    }
  };

  // Offer側でネゴシエーションが必要になったときの処理
  peer.onnegotiationneeded = async () => {
    try {
      if (isOffer) {
        if (negotiationneededCounter === 0) {
          console.log("Creating offer");
          const offer = await peer.createOffer();
          console.log("createOffer() success in promise");
          await peer.setLocalDescription(offer);
          console.log("setLocalDescription() success in promise");
          sendSdp(peer.localDescription);
          negotiationneededCounter++;
        }
      }
    } catch (err) {
      console.error("Error during negotiation: ", err);
    }
  };

  // データチャネルの作成 (オプション)
  if (isOffer) {
    const dataChannel = peer.createDataChannel("chat");
    setupDataChannel(dataChannel);
  } else {
    peer.ondatachannel = (event) => {
      setupDataChannel(event.channel);
    };
  }

  return peer;
}

function setupDataChannel(dataChannel: RTCDataChannel) {
  dataChannel.onopen = () => {
    console.log("Data channel is open");
  };
  dataChannel.onmessage = (event) => {
    console.log("Received message:", event.data);
  };
}

function sendSdp(sessionDescription: RTCSessionDescription | null) {
  console.log("Sending SDP", sessionDescription);
  const textForSendSdp = textForSendSdpRef.current;
  if (textForSendSdp && sessionDescription) {
    textForSendSdp.value = sessionDescription.sdp;
    const message = JSON.stringify(sessionDescription);
    ws.send(message);
    console.log("SDP sent:", message);
  }
}

// Connectボタンが押されたらWebRTCのOffer処理を開始
export function connect(remoteVideoRef: RefObject<HTMLVideoElement>) {
  if (!peerConnection) {
    console.log("Creating new peer connection");
    peerConnection = prepareNewConnection(true, remoteVideoRef);
    if (localStream) {
      localStream.getTracks().forEach((track) => {
        peerConnection!.addTrack(track, localStream);
      });
    } else {
      console.error("Local stream is not available");
    }
    // 明示的にOfferを作成して送信
    createAndSendOffer();
  } else {
    console.warn("Peer connection already exists");
  }
}

async function createAndSendOffer() {
  if (!peerConnection) {
    console.error("PeerConnection does not exist!");
    return;
  }
  try {
    const offer = await peerConnection.createOffer();
    await peerConnection.setLocalDescription(offer);
    console.log("Offer created:", offer);
    sendSdp(peerConnection.localDescription);
  } catch (err) {
    console.error("Error creating offer:", err);
  }
}

// Answer SDPを生成する
async function makeAnswer() {
  if (!peerConnection) {
    console.error("PeerConnection does not exist!");
    return;
  }
  try {
    const answer = await peerConnection.createAnswer();
    await peerConnection.setLocalDescription(answer);
    console.log("Answer created:", answer);
    sendSdp(peerConnection.localDescription);
  } catch (err) {
    console.error("Error creating answer:", err);
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
    console.warn("PeerConnection already exists, closing existing connection");
    peerConnection.close();
  }
  peerConnection = prepareNewConnection(false, remoteVideoRef);
  try {
    await peerConnection.setRemoteDescription(sessionDescription);
    console.log("Remote description set successfully");
    makeAnswer();
  } catch (err) {
    console.error("Error setting remote description:", err);
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
