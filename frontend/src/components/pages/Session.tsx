import { HalfModal } from "../utils/HalfModal";
import { useCallback, useEffect, useRef, useState } from "react";
import { Typography, Box, List, ListItem, Paper, ListItemText, TextField, Button } from "@mui/material";
import HomeLogo from "../../assets/homeLogo";
import { Favorite, Person } from "@mui/icons-material";
import { SessionBottomNavigationTemplate } from "../templates/SessionBottomNavigationTemplate";
import SessionContainer from "../utils/SessionContainer";
import api from "../../services/api";
const users = [
  { name: "User1", icon: <Person />, description: "英語" },
  { name: "User2", icon: <Favorite />, description: "苗字" },
];

const theme = "好きな言葉";

const STUN_SERVERS = {
  iceServers: [
    { urls: "stun:stun.l.google.com:19302" },
    { urls: "stun:stun1.l.google.com:19302" },
  ],
};

export const Session = () => {
  const [memoOpen, setMemoOpen] = useState(false);
  const [assistantOpen, setAssistantOpen] = useState(false);
  const [messages, setMessages] = useState<string[]>([]);
  const [inputMessage, setInputMessage] = useState<string>("");
  const [isLoading, setIsLoading] = useState(false); // ローディング状態を管理

  const handleMemoClose = () => setMemoOpen(false);
  const handleAssistantClose = () => setAssistantOpen(false);

  const handleSendMessage = async () => {
    if (inputMessage.trim() === "") return;

    setIsLoading(true); // ローディング状態を開始
    setMessages([...messages, `You: ${inputMessage}`]); // ユーザーのメッセージを表示
    const userMessage = inputMessage;
    setInputMessage(""); // 送信後に入力フィールドをクリア

    try {
      // サーバーにリクエストを送信
      const response = await api.post("/chat/ask", {
        content: userMessage,
      });

      // レスポンスデータを確認
      console.log("Response data:", response.data); // ここでレスポンスデータを確認

      // 応答メッセージを表示
      if (response.data && response.data.choices && response.data.choices[0].message.content) {
        const assistantMessage = response.data.choices[0].message.content;
        setMessages((prevMessages) => [...prevMessages, `Assistant: ${assistantMessage}`]);
      } else {
        setMessages((prevMessages) => [...prevMessages, "Error: Invalid response format"]);
      }
    } catch (error) {
      setMessages((prevMessages) => [...prevMessages, "Error: Failed to get response"]);
    } finally {
      setIsLoading(false); // ローディング状態を終了
    }
  };

  const memo = "I'm Hanako. Please call me Hanako.";

  // WebRTC関連の処理
  // webbbb vimジャンプ用
  const host = "10.70.174.101"
  const WEBSOCKET_URL = "ws://" + host + ":8081/ws";
  let isOffer = false;
  const [isOfferState, setIsOfferState] = useState<boolean>(false);

  const [isConnected, setIsConnected] = useState<boolean>(false);
  const [isInCall, setIsInCall] = useState<boolean>(false);
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
  }, []);


  useEffect(() => {
    connectToSignalingServer(); // コンポーネントがマウントされたときにシグナリングサーバーに接続する
    //setTimeout(() => startCall(false), 3000); // 3秒後に通話開始
    console.log("通話開始")
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

      const token = localStorage.getItem("token"); // Assuming you store the token in localStorage
      const authMessage = {
        type: 'Authorization',
        token: `Bearer ${token}`
      };
      ws.send(JSON.stringify(authMessage));
      setIsConnected(true);
    };

    ws.onmessage = async (event: MessageEvent) => {
      const message: {
        type: string;
        offer?: RTCSessionDescriptionInit;
        answer?: RTCSessionDescriptionInit;
        candidate?: RTCIceCandidateInit;
        isOffer?: boolean;
      } = JSON.parse(event.data);

      if (message.type == "callType") {
        if (message.isOffer != undefined) {
          isOffer = message.isOffer;
          setIsOfferState(message.isOffer);
          startCall();
        }
      } else if (message.type == "offer" && !isOffer) {
        setTimeout(startCall, 1000);
      }

      if (!peerConnectionRef.current) {
        console.warn("Received message but peer connection is not established");
        return;
      }
      console.log("on message", message)

      console.log("Received message type:", message.type);
      try {
        if (message.type === "offer" && message.offer) {
          await handleOffer(message.offer);
        } else if (message.type === "answer" && message.answer) {
          await peerConnectionRef.current.setRemoteDescription(
            new RTCSessionDescription(message.answer)
          );
          //startCall(true);
        } else if (message.type === "ice-candidate" && message.candidate) {
          await peerConnectionRef.current.addIceCandidate(
            new RTCIceCandidate(message.candidate)
          );
        } else if (message.type === "callType" && message.isOffer) {
          isOffer = message.isOffer;
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
    console.log("start call")
    const pc = createPeerConnection();
    peerConnectionRef.current = pc;

    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      localStreamRef.current = stream;
      stream.getTracks().forEach((track) => pc.addTrack(track, stream));

      const offer = await pc.createOffer();
      await pc.setLocalDescription(offer);

      if (websocketRef.current) {
        websocketRef.current.send(JSON.stringify({
          type: "offer",
          offer: pc.localDescription,
        }));
      }

      setIsInCall(true);
    } catch (error) {
      console.error("Error starting call:", error);
      cleanupResources();
    }
  };

  const handleOffer = async (offer: RTCSessionDescriptionInit): Promise<void> => {
    console.log("handle offer")
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
        websocketRef.current.send(JSON.stringify({
          type: "answer",
          answer: pc.localDescription,
        }));
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
    <SessionBottomNavigationTemplate value="other" isMute={false} setMemoOpen={setMemoOpen} setAssistantOpen={setAssistantOpen}>
      <SessionContainer theme={theme} users={users} />
      <HalfModal open={memoOpen} handleClose={handleMemoClose} title="持ち込みメモ">
        <Typography variant="body1">{memo}</Typography>
      </HalfModal>

      <HalfModal open={assistantOpen} handleClose={handleAssistantClose} title="アシスタント">
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
            {messages.map((message, index) => (
              <ListItem key={index} sx={{ justifyContent: message.startsWith("You:") ? "flex-end" : "flex-start" }}>
                <Paper
                  sx={{
                    padding: "5px",
                    backgroundColor: message.startsWith("You:") ? "#f0f0f0" : "background.default",
                    maxWidth: "60%",
                    wordWrap: "break-word",
                  }}
                >
                  <ListItemText primary={message} />
                </Paper>
              </ListItem>
            ))}
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
            disabled={isLoading} // ローディング中は入力を無効化
          />
          <Button variant="contained" color="primary" onClick={handleSendMessage} disabled={isLoading}>
            {isLoading ? "送信中..." : "送信"}
          </Button>
        </Box>
      </HalfModal>
      <audio ref={remoteAudioRef} autoPlay />
      <p>テスト用: isConnected={isConnected ? "true" : "false"}</p>
      <p>テスト用: isInCall={isInCall ? "true" : "false"}</p>
      <p>テスト用: isOffer={isOfferState ? "true" : "false"}</p>

      <button
        onClick={isInCall ? endCall : () => startCall()}
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
    </SessionBottomNavigationTemplate>
  );
};

export default Session;
