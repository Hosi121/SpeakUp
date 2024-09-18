import React, { useState, useEffect, useRef } from "react";

const WEBSOCKET_URL = "ws://localhost:8080/ws";

const VoiceChat: React.FC = () => {
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const [isInCall, setIsInCall] = useState<boolean>(false);
  const [peerConnection, setPeerConnection] =
    useState<RTCPeerConnection | null>(null);
  const websocketRef = useRef<WebSocket | null>(null);
  const remoteAudioRef = useRef<HTMLAudioElement>(null);
  const localStreamRef = useRef<MediaStream | null>(null);

  useEffect(() => {
    return () => {
      if (websocketRef.current) {
        websocketRef.current.close();
      }
      if (localStreamRef.current) {
        localStreamRef.current.getTracks().forEach((track) => track.stop());
      }
    };
  }, []);

  const connectToSignalingServer = (): void => {
    const ws = new WebSocket(WEBSOCKET_URL);

    ws.onopen = () => {
      console.log("Connected to signaling server");
      setIsConnected(true);
    };

    ws.onmessage = async (event: MessageEvent) => {
      const message: {
        type: string;
        offer?: RTCSessionDescriptionInit;
        answer?: RTCSessionDescriptionInit;
        candidate?: RTCIceCandidateInit;
      } = JSON.parse(event.data);

      if (message.type === "offer" && message.offer) {
        await handleOffer(message.offer);
      } else if (
        message.type === "answer" &&
        message.answer &&
        peerConnection
      ) {
        await peerConnection.setRemoteDescription(
          new RTCSessionDescription(message.answer)
        );
      } else if (
        message.type === "ice-candidate" &&
        message.candidate &&
        peerConnection
      ) {
        await peerConnection.addIceCandidate(
          new RTCIceCandidate(message.candidate)
        );
      }
    };

    ws.onclose = () => {
      console.log("Disconnected from signaling server");
      setIsConnected(false);
    };

    websocketRef.current = ws;
  };

  const startCall = async (): Promise<void> => {
    const pc = new RTCPeerConnection();

    pc.onicecandidate = (event: RTCPeerConnectionIceEvent) => {
      if (event.candidate && websocketRef.current) {
        websocketRef.current.send(
          JSON.stringify({
            type: "ice-candidate",
            candidate: event.candidate,
          })
        );
      }
    };

    pc.ontrack = (event: RTCTrackEvent) => {
      if (remoteAudioRef.current && event.streams[0]) {
        remoteAudioRef.current.srcObject = event.streams[0];
      }
    };

    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      localStreamRef.current = stream;
      stream.getTracks().forEach((track) => pc.addTrack(track, stream));

      const offer = await pc.createOffer();
      await pc.setLocalDescription(offer);

      if (websocketRef.current) {
        websocketRef.current.send(
          JSON.stringify({
            type: "offer",
            offer: pc.localDescription,
          })
        );
      }

      setPeerConnection(pc);
      setIsInCall(true);
    } catch (error) {
      console.error("Error starting call:", error);
    }
  };

  const handleOffer = async (
    offer: RTCSessionDescriptionInit
  ): Promise<void> => {
    const pc = new RTCPeerConnection();

    pc.onicecandidate = (event: RTCPeerConnectionIceEvent) => {
      if (event.candidate && websocketRef.current) {
        websocketRef.current.send(
          JSON.stringify({
            type: "ice-candidate",
            candidate: event.candidate,
          })
        );
      }
    };

    pc.ontrack = (event: RTCTrackEvent) => {
      if (remoteAudioRef.current && event.streams[0]) {
        remoteAudioRef.current.srcObject = event.streams[0];
      }
    };

    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      localStreamRef.current = stream;
      stream.getTracks().forEach((track) => pc.addTrack(track, stream));

      await pc.setRemoteDescription(new RTCSessionDescription(offer));
      const answer = await pc.createAnswer();
      await pc.setLocalDescription(answer);

      if (websocketRef.current) {
        websocketRef.current.send(
          JSON.stringify({
            type: "answer",
            answer: pc.localDescription,
          })
        );
      }

      setPeerConnection(pc);
      setIsInCall(true);
    } catch (error) {
      console.error("Error handling offer:", error);
    }
  };

  const endCall = (): void => {
    if (peerConnection) {
      peerConnection.close();
      setPeerConnection(null);
    }
    if (localStreamRef.current) {
      localStreamRef.current.getTracks().forEach((track) => track.stop());
    }
    if (remoteAudioRef.current) {
      remoteAudioRef.current.srcObject = null;
    }
    setIsInCall(false);
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        gap: "1rem",
      }}
    >
      <button
        onClick={connectToSignalingServer}
        disabled={isConnected}
        style={{
          padding: "10px 20px",
          fontSize: "16px",
          backgroundColor: isConnected ? "#ccc" : "#007bff",
          color: "white",
          border: "none",
          borderRadius: "5px",
          cursor: isConnected ? "default" : "pointer",
        }}
      >
        {isConnected ? "Connected to Server" : "Connect to Server"}
      </button>
      <button
        onClick={isInCall ? endCall : startCall}
        disabled={!isConnected}
        style={{
          padding: "10px 20px",
          fontSize: "16px",
          backgroundColor: !isConnected
            ? "#ccc"
            : isInCall
            ? "#dc3545"
            : "#28a745",
          color: "white",
          border: "none",
          borderRadius: "5px",
          cursor: !isConnected ? "default" : "pointer",
        }}
      >
        {isInCall ? "End Call" : "Start Call"}
      </button>
      <audio ref={remoteAudioRef} autoPlay />
    </div>
  );
};

export default VoiceChat;
