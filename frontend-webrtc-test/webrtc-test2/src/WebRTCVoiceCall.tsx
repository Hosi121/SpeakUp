import React, { useEffect, useRef, useState } from "react";

interface WebSocketMessage {
  type: "offer" | "answer" | "ice_candidate" | "end_call";
  offerId?: number;
  offer?: RTCSessionDescriptionInit;
  answer?: RTCSessionDescriptionInit;
  sdp?: string;
  candidate?: RTCIceCandidate;
}

const WebRTCVoiceCall: React.FC = () => {
  const [isOffer, setIsOffer] = useState<boolean>(false);
  const [isCalling, setIsCalling] = useState<boolean>(false);
  const [localStream, setLocalStream] = useState<MediaStream | null>(null);
  const peerConnection = useRef<RTCPeerConnection | null>(null);
  const websocket = useRef<WebSocket | null>(null);
  const remoteAudioRef = useRef<HTMLAudioElement>(null);

  useEffect(() => {
    // WebSocketの接続
    websocket.current = new WebSocket("ws://localhost:8080/ws");

    websocket.current.onmessage = (event: MessageEvent) => {
      const data: WebSocketMessage = JSON.parse(event.data);
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

    // コンポーネントのアンマウント時にクリーンアップ
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

  const sendOffer = (offer: RTCSessionDescriptionInit): void => {
    const offerId = Math.floor(Math.random() * 1000000);
    if (websocket.current) {
      websocket.current.send(
        JSON.stringify({
          type: "offer",
          offerId: offerId,
          offer: offer,
          sdp: offer.sdp,
        })
      );
    }
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
      }
    } catch (error) {
      console.error("Error handling received offer:", error);
    }
  };

  const sendAnswer = (
    answer: RTCSessionDescriptionInit,
    offerId: number
  ): void => {
    if (websocket.current) {
      websocket.current.send(
        JSON.stringify({
          type: "answer",
          offerId: offerId,
          answer: answer,
          sdp: answer.sdp,
        })
      );
    }
  };

  const handleReceivedAnswer = async (
    data: WebSocketMessage
  ): Promise<void> => {
    try {
      if (data.answer && peerConnection.current) {
        await peerConnection.current.setRemoteDescription(
          new RTCSessionDescription(data.answer)
        );
      }
    } catch (error) {
      console.error("Error handling received answer:", error);
    }
  };

  const handleIceCandidate = async (data: WebSocketMessage): Promise<void> => {
    try {
      if (data.candidate && peerConnection.current) {
        await peerConnection.current.addIceCandidate(
          new RTCIceCandidate(data.candidate)
        );
      }
    } catch (error) {
      console.error("Error handling ICE candidate:", error);
    }
  };

  const sendIceCandidate = (candidate: RTCIceCandidate): void => {
    if (websocket.current) {
      websocket.current.send(
        JSON.stringify({
          type: "ice_candidate",
          candidate: candidate,
        })
      );
    }
  };

  const endCall = (): void => {
    if (websocket.current) {
      websocket.current.send(JSON.stringify({ type: "end_call" }));
    }
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
  };

  return (
    <div>
      {!isCalling ? (
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
