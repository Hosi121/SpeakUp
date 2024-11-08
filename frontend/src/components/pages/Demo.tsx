import React, { useState, useEffect, useRef, useCallback } from "react";
import { SessionBottomNavigationTemplate } from "../templates/SessionBottomNavigationTemplate";
import { HalfModal } from "../utils/HalfModal";
import TabContext from "@mui/lab/TabContext";
import Box from "@mui/material/Box";
import TabList from "@mui/lab/TabList";
import Tab from "@mui/material/Tab";
import TabPanel from "@mui/lab/TabPanel";
import Typography from "@mui/material/Typography";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import HomeLogo from "../../assets/homeLogo";
import Paper from "@mui/material/Paper";
import ListItemText from "@mui/material/ListItemText";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { AudioVolumeAnalyzer } from "../utils/AudioVolumeAnalyzer";
import SessionContainer from "../utils/SessionContainer";
import { Person } from "@mui/icons-material";

const STUN_SERVERS = {
  iceServers: [
    { urls: "stun:stun.l.google.com:19302" },
    { urls: "stun:stun1.l.google.com:19302" },
  ],
};

type HashedId = { hashedId: number };
type SpeakUpLocalStorage = {
  hashedId: number;
  signalingIp: string;
}

export const Demo: React.FC = () => {
  const [developDisplay, setDevelopDisplay] = useState('block');
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const [isInCall, setIsInCall] = useState<boolean>(false);
  const peerConnectionRef = useRef<RTCPeerConnection | null>(null);
  const websocketRef = useRef<WebSocket | null>(null);
  const remoteAudioRef = useRef<HTMLAudioElement>(null);
  const localStreamRef = useRef<MediaStream | null>(null);

  const SPEAKUP_KEY = 'speakupdemo';
  const jsonData = localStorage.getItem(SPEAKUP_KEY) ?? '{"hashedId": 0, "signalingIp": "192.168.1.42"}';
  const speakupStorage = JSON.parse(jsonData) as SpeakUpLocalStorage;
  const [hashedId, setHashedId] = useState(speakupStorage.hashedId);
  const [signalingIp, setSignalingIp] = useState(speakupStorage.signalingIp);
  const saveSpeakupStorage = () => {
    const newStorageData = {
      hashedId: hashedId,
      signalingIp: signalingIp,
    };
    localStorage.setItem(SPEAKUP_KEY, JSON.stringify(newStorageData));
  }
  const WEBSOCKET_URL = `ws://${signalingIp}:8083/ws`;
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
        opponentVolumeAnalyzerRef.current = new AudioVolumeAnalyzer(setIsOpponentSpeak);
        opponentVolumeAnalyzerRef.current.start(event.streams[0]);
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
      volumeAnalyzerRef.current = new AudioVolumeAnalyzer(setisSpeak);
      volumeAnalyzerRef.current.start(stream);
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

  const [isMuted, setIsMuted] = useState<boolean>(false);
  const toggleMute = (): void => {
    setIsMuted((prev) => !prev);
    if (localStreamRef.current) {
      const audioTrack = localStreamRef.current.getAudioTracks()[0];
      if (audioTrack) {
        audioTrack.enabled = !audioTrack.enabled;
        setIsMuted(!audioTrack.enabled);
      }
    }
  };
  // visualize speaker
  const volumeAnalyzerRef = useRef<AudioVolumeAnalyzer | null>(null);
  const opponentVolumeAnalyzerRef = useRef<AudioVolumeAnalyzer | null>(null);
  const [isSpeak, setisSpeak] = useState(false);
  const [isOpponentSpeak, setIsOpponentSpeak] = useState(false);

  // userInfo
  const initialUserCardInfo = { name: "", icon: <Person /> };
  const userCardInfo = initialUserCardInfo;
  const opponentUserCardInfo = initialUserCardInfo;


  const [memoOpen, setMemoOpen] = useState(false);
  const [assistantOpen, setAssistantOpen] = useState(false);
  const handleMemoClose = () => setMemoOpen(false);
  const handleAssistantClose = () => setAssistantOpen(false);
  const [value, setValue] = useState("1");
  const handleChange = (_: React.SyntheticEvent, newValue: string) => setValue(newValue);
  const memo1 = ""; // もともとusestate
  const memo2 = "";
  const [inputMessage, setInputMessage] = useState<string>("");

  return (
    <SessionBottomNavigationTemplate
      value="other"
      isMute={isMuted}
      toggleMute={toggleMute}
      setMemoOpen={setMemoOpen}
      setAssistantOpen={setAssistantOpen}
      onPriorityHighClick={() => { }}
    >
      <SessionContainer
        theme={"好きな言葉"}
        users={[userCardInfo, opponentUserCardInfo]}
        isSpeak={isSpeak && !isMuted}
        isOpponentSpeak={isOpponentSpeak}
      />
      {/* ここから下126行は無視する */}
      <HalfModal open={memoOpen} handleClose={handleMemoClose} title="">
        <TabContext value={value}>
          <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
            <TabList
              onChange={handleChange}
              sx={{ display: "grid", placeContent: "center" }}
            >
              <Tab label=" 持ち込みメモ" value="1" />
              <Tab label="ワードリスト" value="2" />
            </TabList>
          </Box>
          <TabPanel value="1">
            <Typography variant="body1">{memo1}</Typography>
          </TabPanel>
          <TabPanel value="2">
            <Typography variant="body1">{memo2}</Typography>
          </TabPanel>
        </TabContext>
      </HalfModal>
      <HalfModal
        open={assistantOpen}
        handleClose={handleAssistantClose}
        title="アシスタント"
      >
        <Box sx={{ overflow: "auto", pt: 1, pb: 1, maxHeight: "30vh" }}>
          <List>
            {/* アシスタントからの初期メッセージ */}
            <ListItem sx={{ justifyContent: "flex-start" }}>
              <Box
                sx={{
                  width: "40px",
                  height: "40px",
                  borderRadius: "50%",
                  backgroundColor: "secondary.main",
                  mr: 2,
                  display: "grid",
                  placeContent: "center",
                }}
              >
                <HomeLogo style={{ width: "70%", height: "fit-content" }} />
              </Box>
              <Paper
                sx={{
                  padding: "5px",
                  backgroundColor: "background.default",
                  maxWidth: "60%",
                  wordWrap: "break-word",
                }}
              >
                <ListItemText primary="何かお困りですか？" />
              </Paper>
            </ListItem>

            {/* メッセージのリスト */}
            {/* 
            {messages.map((message, index) => (
              <ListItem
                key={index}
                sx={{
                  justifyContent: message.startsWith("You:")
                    ? "flex-end"
                    : "flex-start",
                }}
              >
                <Paper
                  sx={{
                    padding: "5px",
                    backgroundColor: message.startsWith("You:")
                      ? "#f0f0f0"
                      : "background.default",
                    maxWidth: "60%",
                    wordWrap: "break-word",
                  }}
                >
                  <ListItemText primary={message} />
                </Paper>
              </ListItem>
            ))}
            */}
          </List>
        </Box>

        {/* メッセージ入力欄と送信ボタン */}
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            pb: 2,
            position: "fixed",
            bottom: 0,
            backgroundColor: "secondary.main",
          }}
        >
          <TextField
            variant="outlined"
            placeholder="メッセージを入力"
            fullWidth
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
            sx={{ mr: 2 }}
            InputProps={{
              style: {
                height: "40px",
              },
            }}
            disabled={false}
          />
          <Button
            variant="contained"
            color="primary"
            onClick={() => { }}
            disabled={false}
          >
            {"送信"}
          </Button>
        </Box>
      </HalfModal>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: "1rem",
        }}
      >
        {/* ここから上126行は無視する */}
        <Box sx={{ display: developDisplay }}>
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
          <div>
            <p>
              ip: <input value={signalingIp} onChange={(e) => setSignalingIp(e.target.value)} />
            </p>
            <p>
              Hashed ID:
              <input onChange={e => handleSetHashId(e)} value={hashedId} />
            </p>
            <Button onClick={saveSpeakupStorage}>Save</Button>
          </div>
          <Button onClick={() => setDevelopDisplay('none')}>to demo</Button>
        </Box>
      </div>
    </SessionBottomNavigationTemplate>
  );
};
