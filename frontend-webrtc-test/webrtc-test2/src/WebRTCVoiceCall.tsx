import React, { useEffect, useRef, useState } from "react";

interface WebSocketMessage {
  type: "offer" | "answer" | "ice_candidate" | "end_call";
  offerId?: number;
  offer?: RTCSessionDescriptionInit;
  answer?: RTCSessionDescriptionInit;
  sdp?: string;
  candidate?: RTCIceCandidateInit;
}

const WebRTCVoiceCall: React.FC = () => {
  const [isOffer, setIsOffer] = useState<boolean>(false);
  const [isCalling, setIsCalling] = useState<boolean>(false);
  const [localStream, setLocalStream] = useState<MediaStream | null>(null);
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const peerConnection = useRef<RTCPeerConnection | null>(null);
  const websocket = useRef<WebSocket | null>(null);
  const remoteAudioRef = useRef<HTMLAudioElement>(null);
  const iceCandidatesQueue = useRef<RTCIceCandidateInit[]>([]);

  useEffect(() => {
    websocket.current = new WebSocket("ws://10.70.174.101:8080/ws");

    websocket.current.onopen = () => {
      console.log("WebSocket connection established");
      setIsConnected(true);
    };

    websocket.current.onclose = () => {
      console.log("WebSocket connection closed");
      setIsConnected(false);
    };

    websocket.current.onerror = (error) => {
      console.error("WebSocket error:", error);
      setIsConnected(false);
    };

    websocket.current.onmessage = (event: MessageEvent) => {
      const data: WebSocketMessage = JSON.parse(event.data);
      console.log("Received message:", data.type);
      switch (data.type) {
        case "offer":
          handleReceivedOffer(data);
          break;
        case "answer":
          handleReceivedAnswer(data);
          break;
        case "ice_candidate":
          handleIceCandidate(data);
          break;
        case "end_call":
          handleEndCall();
          break;
      }
    };

    return () => {
      cleanupCall();
    };
  }, []);

  const createPeerConnection = (): void => {
    peerConnection.current = new RTCPeerConnection({
      iceServers: [
        { urls: "stun:stun.l.google.com:19302" },
        { urls: "stun:stun1.l.google.com:19302" },
        { urls: "stun:stun2.l.google.com:19302" },
      ],
    });

    peerConnection.current.onicecandidate = (
      event: RTCPeerConnectionIceEvent
    ) => {
      if (event.candidate) {
        sendIceCandidate(event.candidate);
      }
    };

    peerConnection.current.ontrack = (event: RTCTrackEvent) => {
      console.log("Received remote track");
      if (remoteAudioRef.current) {
        remoteAudioRef.current.srcObject = event.streams[0];
      }
    };

    peerConnection.current.oniceconnectionstatechange = () => {
      console.log(
        "ICE connection state:",
        peerConnection.current?.iceConnectionState
      );
    };

    peerConnection.current.onsignalingstatechange = () => {
      console.log("Signaling state:", peerConnection.current?.signalingState);
    };
  };

  const startCall = async (): Promise<void> => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      setLocalStream(stream);
      setIsCalling(true);
      setIsOffer(true);
      createPeerConnection();
      stream
        .getTracks()
        .forEach((track) => peerConnection.current?.addTrack(track, stream));
      const offer = await peerConnection.current?.createOffer();
      await peerConnection.current?.setLocalDescription(offer);
      if (offer) {
        sendOffer(offer);
      }
    } catch (error) {
      console.error("Error starting the call:", error);
    }
  };

  const sendMessage = (message: any): void => {
    if (websocket.current && websocket.current.readyState === WebSocket.OPEN) {
      websocket.current.send(JSON.stringify(message));
    } else {
      console.error("WebSocket is not open. Unable to send message.");
    }
  };

  const sendOffer = (offer: RTCSessionDescriptionInit): void => {
    const offerId = Math.floor(Math.random() * 1000000);
    sendMessage({
      type: "offer",
      offerId: offerId,
      offer: offer,
      sdp: offer.sdp,
    });
  };

  const handleReceivedOffer = async (data: WebSocketMessage): Promise<void> => {
    try {
      createPeerConnection();
      if (data.offer && peerConnection.current) {
        await peerConnection.current.setRemoteDescription(
          new RTCSessionDescription(data.offer)
        );
        const stream = await navigator.mediaDevices.getUserMedia({
          audio: true,
        });
        setLocalStream(stream);
        stream
          .getTracks()
          .forEach((track) => peerConnection.current?.addTrack(track, stream));
        const answer = await peerConnection.current.createAnswer();
        await peerConnection.current.setLocalDescription(answer);
        if (data.offerId !== undefined) {
          sendAnswer(answer, data.offerId);
        }
        setIsCalling(true);
        processIceCandidateQueue();
      }
    } catch (error) {
      console.error("Error handling received offer:", error);
    }
  };

  const sendAnswer = (
    answer: RTCSessionDescriptionInit,
    offerId: number
  ): void => {
    sendMessage({
      type: "answer",
      offerId: offerId,
      answer: answer,
      sdp: answer.sdp,
    });
  };

  const handleReceivedAnswer = async (
    data: WebSocketMessage
  ): Promise<void> => {
    try {
      if (data.answer && peerConnection.current) {
        await peerConnection.current.setRemoteDescription(
          new RTCSessionDescription(data.answer)
        );
        processIceCandidateQueue();
      }
    } catch (error) {
      console.error("Error handling received answer:", error);
    }
  };

  const handleIceCandidate = (data: WebSocketMessage): void => {
    if (data.candidate) {
      const iceCandidate = new RTCIceCandidate(data.candidate);
      if (
        peerConnection.current &&
        peerConnection.current.remoteDescription &&
        peerConnection.current.remoteDescription.type
      ) {
        peerConnection.current
          .addIceCandidate(iceCandidate)
          .catch((e) =>
            console.error("Error adding received ice candidate:", e)
          );
      } else {
        iceCandidatesQueue.current.push(data.candidate);
      }
    }
  };

  const processIceCandidateQueue = (): void => {
    if (peerConnection.current && peerConnection.current.remoteDescription) {
      iceCandidatesQueue.current.forEach((candidate) => {
        peerConnection.current
          ?.addIceCandidate(new RTCIceCandidate(candidate))
          .catch((e) => console.error("Error adding queued ice candidate:", e));
      });
      iceCandidatesQueue.current = [];
    }
  };

  const sendIceCandidate = (candidate: RTCIceCandidate): void => {
    sendMessage({
      type: "ice_candidate",
      candidate: candidate,
    });
  };

  const endCall = (): void => {
    sendMessage({ type: "end_call" });
    cleanupCall();
  };

  const handleEndCall = (): void => {
    cleanupCall();
  };

  const cleanupCall = (): void => {
    if (peerConnection.current) {
      peerConnection.current.close();
      peerConnection.current = null;
    }
    if (localStream) {
      localStream.getTracks().forEach((track) => track.stop());
      setLocalStream(null);
    }
    setIsCalling(false);
    setIsOffer(false);
    if (remoteAudioRef.current) {
      remoteAudioRef.current.srcObject = null;
    }
    iceCandidatesQueue.current = [];
  };

  return (
    <div>
      {!isConnected ? (
        <p>Connecting to server...</p>
      ) : !isCalling ? (
        <button onClick={startCall}>Start Call</button>
      ) : (
        <div>
          <p>{isOffer ? "Calling..." : "In Call"}</p>
          <button onClick={endCall}>End Call</button>
        </div>
      )}
      <audio ref={remoteAudioRef} autoPlay />
    </div>
  );
};

export default WebRTCVoiceCall;
