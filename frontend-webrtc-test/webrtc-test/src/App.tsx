import "./App.css";
import { connect, hangUp, onSdpText, startVideo } from "./webrtc";
import styles from "./App.module.css";
import React, { useRef } from "react";

function App() {
  navigator.mediaDevices
    .getUserMedia({ video: true, audio: true })
    .then(function (stream) {
      // success
      console.log(stream);
    })
    .catch(function (error) {
      // error
      console.log(error);
      return;
    });

  const localVideoRef = useRef(null);
  const remoteVideoRef = useRef(null);

  return (
    <div>
      WebRTC test
      <br />
      <button type="button" onClick={() => startVideo(localVideoRef)}>
        Start Video
      </button>
      <button type="button" onClick={() => connect()}>
        Connect
      </button>
      <button type="button" onClick={() => hangUp()}>
        Hang Up
      </button>
      <div>
        <video
          className={styles.video}
          ref={localVideoRef}
          id="local_video"
          autoPlay
          muted
        ></video>
        <video
          className={styles.video}
          ref={remoteVideoRef}
          id="remote_video"
          autoPlay
        ></video>
      </div>
      <p>
        SDP to send:
        <br />
        <textarea className={styles.textArea} id="text_for_send_sdp" readOnly>
          SDP to send
        </textarea>
      </p>
      <p>
        SDP to receive:
        <button type="button" onClick={() => onSdpText()}>
          Receive remote SDP
        </button>
        <br />
        <textarea
          className={styles.textArea}
          id="text_for_receive_sdp"
        ></textarea>
      </p>
    </div>
  );
}

export default App;
