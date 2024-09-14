import React, { useEffect, useRef, useState } from "react";
import "./App.css";
import { connect, hangUp, onSdpText, startLocalAudioStream } from "./webrtc";
import styles from "./App.module.css";

export let remoteAudioRef: React.RefObject<HTMLAudioElement>;
export let textForSendSdpRef: React.RefObject<HTMLTextAreaElement>;
export let textToReceiveSdpRef: React.RefObject<HTMLTextAreaElement>;

function App() {
  const [mediaDeviceStatus, setMediaDeviceStatus] = useState<
    "checking" | "available" | "unavailable" | "error"
  >("checking");
  const [isCallActive, setIsCallActive] = useState(false);
  remoteAudioRef = useRef<HTMLAudioElement>(null);
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
        console.log("Audio device accessed successfully", stream);
        setMediaDeviceStatus("available");
        stream.getTracks().forEach((track) => track.stop());
      } catch (error) {
        console.error("Error accessing audio device:", error);
        setMediaDeviceStatus("error");
      }
    }

    initMediaDevices();
  }, []);

  const handleStartCall = async () => {
    if (mediaDeviceStatus === "available") {
      await startLocalAudioStream();
      connect(remoteAudioRef);
      setIsCallActive(true);
    } else {
      console.warn("Audio device is not available");
    }
  };

  const handleHangUp = () => {
    if (textForSendSdpRef.current && textToReceiveSdpRef.current) {
      hangUp(textForSendSdpRef.current, textToReceiveSdpRef.current);
      setIsCallActive(false);
    }
  };

  const handleReceiveSdp = () => {
    if (textToReceiveSdpRef.current) {
      onSdpText(textToReceiveSdpRef.current);
    }
  };

  return (
    <div>
      <h1>WebRTC Audio Call</h1>
      {mediaDeviceStatus === "checking" && <p>Checking audio device...</p>}
      {mediaDeviceStatus === "unavailable" && (
        <p>Audio device is not available</p>
      )}
      {mediaDeviceStatus === "error" && <p>Error accessing audio device</p>}
      {mediaDeviceStatus === "available" && (
        <>
          <button
            type="button"
            onClick={handleStartCall}
            disabled={isCallActive}
          >
            Start Call
          </button>
          <button type="button" onClick={handleHangUp} disabled={!isCallActive}>
            Hang Up
          </button>
          <div>
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
