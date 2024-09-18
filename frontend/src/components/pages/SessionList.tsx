import { HalfModal } from "../utils/HalfModal";
import { useState } from "react";
import { Typography } from "@mui/material";
import { Favorite, Person} from "@mui/icons-material";
import { SessionBottomNavigationTemplate } from "../templates/SessionBottomNavigationTemplate";
import SessionContainer from "../utils/SessionContainer";

const users = [
  { name: "User1", icon: <Person />, description: "英語" },
  { name: "User2", icon: <Favorite />, description: "苗字" },
];

const theme = "好きな言葉";

export const Session = () => {
  const [memoOpen, setMemoOpen] = useState(false);
  const [assistantOpen, setAssistantOpen] = useState(false);
  const [messages, setMessages] = useState<string[]>([]);
  const [inputMessage, setInputMessage] = useState<string>("");

  const handleMemoClose = () => setMemoOpen(false);
  const handleAssistantClose = () => setAssistantOpen(false);
  const handleSendMessage = () => {
    if (inputMessage.trim() !== "") {
      setMessages([...messages, inputMessage]);
      setInputMessage(""); // Clear the input field after sending
    }
  };

  const memo = "I'm Hanako. Please call me Hanako.";

  return (
    <SessionBottomNavigationTemplate value="other" isMute={false} setMemoOpen={setMemoOpen} setAssistantOpen={setAssistantOpen}>
      <SessionContainer theme={theme} users={users} />
      <HalfModal open={memoOpen} handleClose={handleMemoClose} title="持ち込みメモ">
        <Typography variant="body1">{memo}</Typography>
      </HalfModal>

      <HalfModal open={assistantOpen} handleClose={handleAssistantClose} title="アシスタント">
        {/* Assistant content here */}
      </HalfModal>
    </SessionBottomNavigationTemplate>
  );
};