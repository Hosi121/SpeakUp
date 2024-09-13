import React, { useEffect, useRef, useState } from "react";
import "./App.css";
import { connect, hangUp, onSdpText, startVideo } from "./webrtc";
import styles from "./App.module.css";

export let remoteVideoRef: React.RefObject<HTMLVideoElement>;
export let textForSendSdpRef: React.RefObject<HTMLTextAreaElement>;
export let textToReceiveSdpRef: React.RefObject<HTMLTextAreaElement>;

function App() {
  const [mediaDeviceStatus, setMediaDeviceStatus] = useState<
    "checking" | "available" | "unavailable" | "error"
  >("checking");
  const localVideoRef = useRef<HTMLVideoElement>(null);
  remoteVideoRef = useRef<HTMLVideoElement>(null);
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
          video: true,
          audio: true,
        });
        console.log("Media devices accessed successfully", stream);
        setMediaDeviceStatus("available");
        // ストリームを停止 (必要に応じて)
        // stream.getTracks().forEach(track => track.stop());
      } catch (error) {
        console.error("Error accessing media devices:", error);
        setMediaDeviceStatus("error");
      }
    }

    initMediaDevices();
  }, []);

  const handleStartVideo = () => {
    if (mediaDeviceStatus === "available") {
      startVideo(localVideoRef);
    } else {
      console.warn("Media devices are not available");
    }
  };

  const handleConnect = () => {
    if (mediaDeviceStatus === "available") {
      connect();
    } else {
      console.warn("Media devices are not available for connection");
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
      <h1>WebRTC test</h1>
      {mediaDeviceStatus === "checking" && <p>Checking media devices...</p>}
      {mediaDeviceStatus === "unavailable" && (
        <p>Media devices are not available</p>
      )}
      {mediaDeviceStatus === "error" && <p>Error accessing media devices</p>}
      {mediaDeviceStatus === "available" && (
        <>
          <button type="button" onClick={handleStartVideo}>
            Start Video
          </button>
          <button type="button" onClick={handleConnect}>
            Connect
          </button>
          <button type="button" onClick={handleHangUp}>
            Hang Up
          </button>
          <div>
            <video
              className={styles.video}
              ref={localVideoRef}
              id="local_video"
              autoPlay
              muted
            />
            <video
              className={styles.video}
              ref={remoteVideoRef}
              id="remote_video"
              autoPlay
            />
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
