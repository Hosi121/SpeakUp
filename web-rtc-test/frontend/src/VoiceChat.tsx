import React, { useState, useEffect, useRef, useCallback } from "react";

const WEBSOCKET_URL = "ws://10.70.174.101:8080/ws";
const STUN_SERVERS = {
  iceServers: [
    { urls: "stun:stun.l.google.com:19302" },
    { urls: "stun:stun1.l.google.com:19302" },
  ],
};

type HashedId = { hashedId: number };

const VoiceChat: React.FC = () => {
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const [isInCall, setIsInCall] = useState<boolean>(false);
  const peerConnectionRef = useRef<RTCPeerConnection | null>(null);
  const websocketRef = useRef<WebSocket | null>(null);
  const remoteAudioRef = useRef<HTMLAudioElement>(null);
  const localStreamRef = useRef<MediaStream | null>(null);

  const [hashedId, setHashedId] = useState(0);
  const handleSetHashId = (e: React.ChangeEvent<HTMLInputElement>) => {
    const id = Number(e.target.value);
    setHashedId(id);
  }

  const cleanupResources = useCallback(() => {
    if (peerConnectionRef.current) {
      peerConnectionRef.current.close();
      peerConnectionRef.current = null;
    }
    if (localStreamRef.current) {
      localStreamRef.current.getTracks().forEach((track) => track.stop());
      localStreamRef.current = null;
    }
    if (remoteAudioRef.current) {
      remoteAudioRef.current.srcObject = null;
    }
    setIsInCall(false);
  }, []);

  useEffect(() => {
    return () => {
      if (websocketRef.current) {
        websocketRef.current.close();
      }
      cleanupResources();
    };
  }, [cleanupResources]);

  const connectToSignalingServer = (): void => {
    const ws = new WebSocket(WEBSOCKET_URL);

    ws.onopen = () => {
      console.log("Connected to signaling server");
      setIsConnected(true);

      const idData: HashedId = { hashedId: hashedId };
      ws.send(JSON.stringify(idData));
    };

    ws.onmessage = async (event: MessageEvent) => {
      const message: {
        type: string;
        offer?: RTCSessionDescriptionInit;
        answer?: RTCSessionDescriptionInit;
        candidate?: RTCIceCandidateInit;
      } = JSON.parse(event.data);

      if (!peerConnectionRef.current) {
        console.warn("Received message but peer connection is not established");
        return;
      }

      try {
        if (message.type === "offer" && message.offer) {
          await handleOffer(message.offer);
        } else if (message.type === "answer" && message.answer) {
          await peerConnectionRef.current.setRemoteDescription(
            new RTCSessionDescription(message.answer)
          );
        } else if (message.type === "ice-candidate" && message.candidate) {
          await peerConnectionRef.current.addIceCandidate(
            new RTCIceCandidate(message.candidate)
          );
        }
      } catch (error) {
        console.error("Error handling WebSocket message:", error);
      }
    };

    ws.onclose = () => {
      console.log("Disconnected from signaling server");
      setIsConnected(false);
      cleanupResources();
    };

    websocketRef.current = ws;
  };

  const createPeerConnection = (): RTCPeerConnection => {
    const pc = new RTCPeerConnection(STUN_SERVERS);

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

    pc.oniceconnectionstatechange = () => {
      console.log("ICE Connection State:", pc.iceConnectionState);
      if (
        pc.iceConnectionState === "disconnected" ||
        pc.iceConnectionState === "failed" ||
        pc.iceConnectionState === "closed"
      ) {
        cleanupResources();
      }
    };

    return pc;
  };

  const startCall = async (): Promise<void> => {
    const pc = createPeerConnection();
    peerConnectionRef.current = pc;

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

      setIsInCall(true);
    } catch (error) {
      console.error("Error starting call:", error);
      cleanupResources();
    }
  };

  const handleOffer = async (
    offer: RTCSessionDescriptionInit
  ): Promise<void> => {
    const pc = createPeerConnection();
    peerConnectionRef.current = pc;

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

      setIsInCall(true);
    } catch (error) {
      console.error("Error handling offer:", error);
      cleanupResources();
    }
  };

  const endCall = (): void => {
    cleanupResources();
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
      <p>
        Hashed ID:
        <input onChange={e => handleSetHashId(e)} />
      </p>
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
