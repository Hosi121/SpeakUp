import React, { useEffect, useRef, useState } from "react";
import "./App.css";
import {
  connect,
  hangUp,
  onSdpText,
  startLocalStream,
  setLocalStreamMute,
  setOffer,
  addIceCandidate,
  setAnswer,
} from "./webrtc";
import styles from "./App.module.css";
import axios from "axios";

export let textForSendSdpRef: React.RefObject<HTMLTextAreaElement>;
export let textToReceiveSdpRef: React.RefObject<HTMLTextAreaElement>;
export let ws: WebSocket;

type UserId = {
  userId: number;
};

type SessionInfo = {
  sessionId: number;
  attribute: string;
};

function App() {
  const [isMuted, setIsMuted] = useState(false);
  const [userId, setUserId] = useState(0);
  const [attribute, setAttribute] = useState("");
  const [sessionId, setSessionId] = useState(0);
  const [mediaDeviceStatus, setMediaDeviceStatus] = useState<
    "checking" | "available" | "unavailable" | "error"
  >("checking");
  const localAudioRef = useRef<HTMLAudioElement>(null);
  const remoteAudioRef = useRef<HTMLAudioElement>(null);
  textForSendSdpRef = useRef<HTMLTextAreaElement>(null);
  textToReceiveSdpRef = useRef<HTMLTextAreaElement>(null);

  useEffect(() => {
    async function initMediaDevices() {
      if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
        console.warn("navigator.mediaDevices.getUserMedia is not supported");
        setMediaDeviceStatus("unavailable");
        return;
      }

      try {
        const stream = await navigator.mediaDevices.getUserMedia({
          audio: true,
        });
        console.log("Audio devices accessed successfully", stream);
        setMediaDeviceStatus("available");
      } catch (error) {
        console.error("Error accessing audio devices:", error);
        setMediaDeviceStatus("error");
      }
    }

    initMediaDevices();
  }, []);

  const handleStartAudio = () => {
    if (mediaDeviceStatus === "available") {
      startLocalStream(localAudioRef);
    } else {
      console.warn("Audio devices are not available");
    }
  };

  const handleConnect = () => {
    if (mediaDeviceStatus === "available") {
      connect();
    } else {
      console.warn("Audio devices are not available for connection");
    }
  };

  const handleHangUp = () => {
    if (textForSendSdpRef.current && textToReceiveSdpRef.current) {
      hangUp(textForSendSdpRef.current, textToReceiveSdpRef.current);
    }
  };

  const handleReceiveSdp = () => {
    if (textToReceiveSdpRef.current) {
      onSdpText(textToReceiveSdpRef.current);
    }
  };

  const handleToggleMute = () => {
    const newState = !isMuted;
    setIsMuted(newState);
    setLocalStreamMute(newState);
  };

  const handleEnterSession = () => {
    // 登録処理
    let attribute = "";
    let sessionId;
    axios
      .post("http://localhost:3001/userId", {
        userId: userId,
      } as UserId)
      .then((res) => {
        console.log(res);
        sessionId = res.data.sessionId;
        attribute = res.data.attribute;
        setSessionId(sessionId);
        setAttribute(attribute);
      });
    // シグナリングサーバへ接続する
    const wsUrl = "ws://10.70.174.103:3001/";
    ws = new WebSocket(wsUrl);
    ws.onopen = () => {
      console.log("ws open()");
      if (attribute === "offer") {
        const data = { userId: userId, sbd: {} };
        ws.send(JSON.stringify(data));
      }
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
          await setOffer(new RTCSessionDescription(message));
          break;
        }
        case "answer": {
          console.log("Received answer ...");
          const textToReceiveSdp = textToReceiveSdpRef.current;
          if (textToReceiveSdp) {
            textToReceiveSdp.value = message.sdp;
          }
          setAnswer(new RTCSessionDescription(message));
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
        case "sbd": {
          console.log("peer is closed ...");
          break;
        }
        default: {
          console.log("Invalid message");
          break;
        }
      }
    };
  };

  return (
    <div>
      <h1>WebRTC Audio-only Test</h1>
      {mediaDeviceStatus === "checking" && <p>Checking audio devices...</p>}
      {mediaDeviceStatus === "unavailable" && (
        <p>Audio devices are not available</p>
      )}
      {mediaDeviceStatus === "error" && <p>Error accessing audio devices</p>}
      {mediaDeviceStatus === "available" && (
        <>
          <button type="button" onClick={handleToggleMute}>
            {isMuted ? "current: Mute" : "current: UnMute"}
          </button>
          <button type="button" onClick={handleStartAudio}>
            Start Audio
          </button>
          <button type="button" onClick={handleConnect}>
            Connect
          </button>
          <button type="button" onClick={handleHangUp}>
            Hang Up
          </button>
          <div>
            <audio ref={localAudioRef} id="local_audio" autoPlay muted />
            <audio ref={remoteAudioRef} id="remote_audio" autoPlay />
          </div>
          <p>
            SDP to send:
            <br />
            <textarea
              className={styles.textArea}
              ref={textForSendSdpRef}
              id="text_for_send_sdp"
              readOnly
              defaultValue="SDP to send"
            />
          </p>
          <p>
            SDP to receive:
            <button type="button" onClick={handleReceiveSdp}>
              Receive remote SDP
            </button>
            <br />
            <textarea
              className={styles.textArea}
              ref={textToReceiveSdpRef}
              id="text_for_receive_sdp"
            />
          </p>
          <p>
            userId: {userId}, sessionId: {sessionId}, attribute: {attribute}
            <br />
            <textarea
              className={styles.textArea}
              onChange={(e) => setUserId(Number(e.target.value))}
            />
            <br />
            <button>セッションに入る</button>
          </p>
        </>
      )}
    </div>
  );
}

export default App;
