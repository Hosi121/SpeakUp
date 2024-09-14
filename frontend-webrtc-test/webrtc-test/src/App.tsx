import React, { useEffect, useRef, useState } from "react";
import "./App.css";
import { connect, hangUp, onSdpText, startLocalStream } from "./webrtc";
import styles from "./App.module.css";

export let textForSendSdpRef: React.RefObject<HTMLTextAreaElement>;
export let textToReceiveSdpRef: React.RefObject<HTMLTextAreaElement>;

function App() {
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
        // ストリームを停止 (必要に応じて)
        // stream.getTracks().forEach(track => track.stop());
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
      connect(remoteAudioRef);
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
        </>
      )}
    </div>
  );
}

export default App;
