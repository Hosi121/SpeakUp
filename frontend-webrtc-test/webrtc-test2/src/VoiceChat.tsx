import React, { useState, useEffect, useRef, useCallback } from "react";

const WEBSOCKET_URL = "ws://10.70.174.101:8080/ws";
const STUN_SERVERS = {
  iceServers: [
    { urls: "stun:stun.l.google.com:19302" },
    { urls: "stun:stun1.l.google.com:19302" },
  ],
};

const VoiceChat: React.FC = () => {
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const [isInCall, setIsInCall] = useState<boolean>(false);
  const [isIncomingCall, setIsIncomingCall] = useState<boolean>(false);
  const peerConnectionRef = useRef<RTCPeerConnection | null>(null);
  const websocketRef = useRef<WebSocket | null>(null);
  const remoteAudioRef = useRef<HTMLAudioElement>(null);
  const localStreamRef = useRef<MediaStream | null>(null);

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
    setIsIncomingCall(false);
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
    };

    ws.onmessage = async (event: MessageEvent) => {
      const message: {
        type: string;
        offer?: RTCSessionDescriptionInit;
        answer?: RTCSessionDescriptionInit;
        candidate?: RTCIceCandidateInit;
      } = JSON.parse(event.data);

      try {
        switch (message.type) {
          case "call_request":
            setIsIncomingCall(true);
            break;
          case "call_accepted":
            await startCall(true);
            break;
          case "offer":
            if (message.offer) await handleOffer(message.offer);
            break;
          case "answer":
            if (message.answer && peerConnectionRef.current) {
              await peerConnectionRef.current.setRemoteDescription(
                new RTCSessionDescription(message.answer)
              );
            }
            break;
          case "ice-candidate":
            if (message.candidate && peerConnectionRef.current) {
              await peerConnectionRef.current.addIceCandidate(
                new RTCIceCandidate(message.candidate)
              );
            }
            break;
          default:
            console.warn("Unknown message type:", message.type);
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

  const startCall = async (isReceiver: boolean = false): Promise<void> => {
    const pc = createPeerConnection();
    peerConnectionRef.current = pc;

    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      localStreamRef.current = stream;
      stream.getTracks().forEach((track) => pc.addTrack(track, stream));

      if (!isReceiver) {
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
      }

      setIsInCall(true);
      setIsIncomingCall(false);
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
    } catch (error) {
      console.error("Error handling offer:", error);
      cleanupResources();
    }
  };

  const requestCall = (): void => {
    if (websocketRef.current) {
      websocketRef.current.send(JSON.stringify({ type: "call_request" }));
    }
  };

  const acceptCall = async (): Promise<void> => {
    if (websocketRef.current) {
      websocketRef.current.send(JSON.stringify({ type: "call_accepted" }));
      await startCall(true);
    }
  };

  const endCall = (): void => {
    cleanupResources();
    if (websocketRef.current) {
      websocketRef.current.send(JSON.stringify({ type: "call_ended" }));
    }
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
      {isIncomingCall ? (
        <button
          onClick={acceptCall}
          style={{
            padding: "10px 20px",
            fontSize: "16px",
            backgroundColor: "#28a745",
            color: "white",
            border: "none",
            borderRadius: "5px",
            cursor: "pointer",
          }}
        >
          Accept Call
        </button>
      ) : isInCall ? (
        <button
          onClick={endCall}
          style={{
            padding: "10px 20px",
            fontSize: "16px",
            backgroundColor: "#dc3545",
            color: "white",
            border: "none",
            borderRadius: "5px",
            cursor: "pointer",
          }}
        >
          End Call
        </button>
      ) : (
        <button
          onClick={requestCall}
          disabled={!isConnected}
          style={{
            padding: "10px 20px",
            fontSize: "16px",
            backgroundColor: !isConnected ? "#ccc" : "#28a745",
            color: "white",
            border: "none",
            borderRadius: "5px",
            cursor: !isConnected ? "default" : "pointer",
          }}
        >
          Start Call
        </button>
      )}
      <audio ref={remoteAudioRef} autoPlay />
    </div>
  );
};

export default VoiceChat;
