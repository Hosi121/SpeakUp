import { Message, Mic, MicOff, PriorityHigh, Search } from "@mui/icons-material";
import { BottomNavigation, BottomNavigationAction, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";

export type sessionBottomNavigation = "wordList" | "assistance" | "topic" | "other";
type sessionBottomNavigationProps = {
  value: sessionBottomNavigation;
  isMute: boolean;
  setMemoOpen: (open: boolean) => void;
  setAssistantOpen: (open: boolean) => void;
};

export const SessionBottomNavigation = ({ value, isMute, setMemoOpen, setAssistantOpen }: sessionBottomNavigationProps) => {
  const navigate = useNavigate();
  const sessionBottomNavigationToIndex = (value: sessionBottomNavigation): number => {
    switch (value) {
      case "wordList":
        return 0;
      case "assistance":
        return 1;
      case "topic":
        return 2;
      default:
        return -1;
    }
  };
  const index = sessionBottomNavigationToIndex(value);
  return (
    <Box>
      <BottomNavigation
        showLabels
        value={index}
        sx={{
          padding: "20px 0",
          height: "70px",
          position: "sticky",
          bottom: 0,
          zIndex: 100,
          backgroundColor: "background.default",
          boxShadow: "0 0 5px rgba(0, 0, 0, 0.6)",
        }}
      >
        <BottomNavigationAction onClick={() => navigate("")} icon={isMute ? <MicOff fontSize="large" /> : <Mic fontSize="large" />} />
        <BottomNavigationAction onClick={() => setMemoOpen(true)} icon={<Message fontSize="large" />} />
        <BottomNavigationAction onClick={() => setAssistantOpen(true)} icon={<Search fontSize="large" />} />
        <BottomNavigationAction onClick={() => navigate("")} icon={<PriorityHigh fontSize="large" />} />
      </BottomNavigation>
    </Box>
  );
};
