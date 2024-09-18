import { RefObject } from "react";
import { textForSendSdpRef, ws } from "./App";

let localStream: MediaStream | null = null;
let peerConnection: RTCPeerConnection | null;
let negotiationneededCounter = 0;
const iceCandidateQueue: RTCIceCandidate[] = [];

export const addIceCandidate = (candidate: RTCIceCandidate) => {
  if (
    peerConnection &&
    peerConnection.remoteDescription &&
    peerConnection.remoteDescription.type
  ) {
    peerConnection
      .addIceCandidate(candidate)
      .then(() => console.log("Added ICE candidate successfully"))
      .catch((err) =>
        console.error("Error adding received ICE candidate", err)
      );
  } else {
    iceCandidateQueue.push(candidate);
    console.log("ICE candidate added to queue");
  }
};

const sendIceCandidate = (candidate: RTCIceCandidate) => {
  console.log("---sending ICE candidate ---");
  const message = JSON.stringify({ type: "candidate", ice: candidate });
  console.log("sending candidate=" + message);
  ws.send(message);
};

export const startLocalStream = async (
  localAudio: RefObject<HTMLAudioElement>
) => {
  try {
    localStream = await navigator.mediaDevices.getUserMedia({
      audio: true,
    });
    if (localAudio.current) {
      localAudio.current.srcObject = localStream;
    }
  } catch (err) {
    console.error("mediaDevice.getUserMedia() error:", err);
    throw err;
  }
};

export const setLocalStreamMute = (isMuted: boolean) => {
  if (localStream) {
    const audioTrack = localStream.getAudioTracks()[0];
    if (audioTrack) {
      audioTrack.enabled = !isMuted;
      console.log(isMuted ? "Muted" : "Unmuted");
    }
  }
  return isMuted;
};

const prepareNewConnection = (isOffer: boolean) => {
  const pc_config = {
    iceServers: [
      { urls: "stun:stun.webrtc.ecl.ntt.com:3478" },
      { urls: "stun:stun.l.google.com:19302" },
    ],
  };
  const peer = new RTCPeerConnection(pc_config);

  peer.ontrack = (evt) => {
    console.log("-- peer.ontrack()", evt.track.kind);
    if (evt.track.kind === "audio") {
      const remoteAudio = new Audio();
      remoteAudio.srcObject = evt.streams[0];
      remoteAudio.play();
      console.log("Remote audio playback started");
    }
  };

  peer.onicecandidate = (evt) => {
    if (evt.candidate) {
      console.log("New ICE candidate: ", evt.candidate);
      sendIceCandidate(evt.candidate);
    } else {
      console.log("ICE gathering completed");
    }
  };

  peer.oniceconnectionstatechange = () => {
    console.log("ICE connection state changed to: ", peer.iceConnectionState);
    switch (peer.iceConnectionState) {
      case "connected":
        console.log("ICE connected successfully");
        break;
      case "failed":
        console.log("ICE connection failed");
        retryConnection();
        break;
      case "disconnected":
        console.log("ICE disconnected");
        retryConnection();
        break;
    }
  };

  const retryConnection = () => {
    if (peerConnection) {
      peerConnection.restartIce();
      if (peerConnection.iceConnectionState === "failed") {
        createAndSendOffer();
      }
    }
  };

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

  if (isOffer) {
    const dataChannel = peer.createDataChannel("chat");
    setupDataChannel(dataChannel);
  } else {
    peer.ondatachannel = (event) => {
      setupDataChannel(event.channel);
    };
  }

  return peer;
};

const setupDataChannel = (dataChannel: RTCDataChannel) => {
  dataChannel.onopen = () => {
    console.log("Data channel is open");
  };
  dataChannel.onmessage = (event) => {
    console.log("Received message:", event.data);
  };
};

const sendSdp = (sessionDescription: RTCSessionDescription | null) => {
  console.log("---sending sdp ---");
  const textForSendSdp = textForSendSdpRef.current;
  if (textForSendSdp && sessionDescription) {
    textForSendSdp.value = sessionDescription.sdp;
  }
  const message = JSON.stringify(sessionDescription);
  console.log("sending SDP=" + message);
  ws.send(message);
};

export const connect = () => {
  if (!peerConnection) {
    console.log("Creating new peer connection");
    peerConnection = prepareNewConnection(true);
    if (localStream) {
      localStream.getTracks().forEach((track) => {
        if (peerConnection && localStream) {
          peerConnection.addTrack(track, localStream);
        }
      });
    } else {
      console.log("No local stream available, connecting without audio");
    }
    createAndSendOffer();
  } else {
    console.warn("Peer connection already exists");
  }
};

const createAndSendOffer = async () => {
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
};

const makeAnswer = async () => {
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
};

export const onSdpText = (textToReceiveSdp: HTMLTextAreaElement | null) => {
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
};

const setRemoteDescription = async (
  sessionDescription: RTCSessionDescription
) => {
  if (!peerConnection) {
    console.error("PeerConnection does not exist!");
    return;
  }
  try {
    await peerConnection.setRemoteDescription(sessionDescription);
    console.log("Remote description set successfully");

    while (iceCandidateQueue.length > 0) {
      const candidate = iceCandidateQueue.shift();
      await peerConnection.addIceCandidate(candidate!);
      console.log("Queued ICE candidate added");
    }

    if (sessionDescription.type === "offer") {
      await makeAnswer();
    }
  } catch (err) {
    console.error("Error setting remote description:", err);
  }
};

export const setOffer = async (sessionDescription: RTCSessionDescription) => {
  if (peerConnection) {
    console.warn("PeerConnection already exists, closing existing connection");
    peerConnection.close();
  }
  peerConnection = prepareNewConnection(false);

  if (localStream) {
    localStream.getTracks().forEach((track) => {
      if (peerConnection && localStream) {
        peerConnection.addTrack(track, localStream);
      }
    });
  } else {
    console.log("No local stream available, connecting without audio");
  }

  await setRemoteDescription(sessionDescription);
};

export const setAnswer = async (sessionDescription: RTCSessionDescription) => {
  if (!peerConnection) {
    console.error("peerConnection NOT exist!");
    return;
  }
  await setRemoteDescription(sessionDescription);
};

export const hangUp = (
  textForSendSdp: HTMLTextAreaElement | null,
  textToReceiveSdp: HTMLTextAreaElement | null
) => {
  if (peerConnection) {
    if (peerConnection.iceConnectionState !== "closed") {
      peerConnection.close();
      peerConnection = null;
      negotiationneededCounter = 0;
      const message = JSON.stringify({ type: "close" });
      console.log("sending close message");
      ws.send(message);

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
};
