import { ReactNode } from "react";
import { SessionBottomNavigation, sessionBottomNavigation } from "../utils/SessionBottomNavigation";
import { Box } from "@mui/material";

type sessionBottomNavigationTemplateProps = {
  value: sessionBottomNavigation;
  isMute: boolean;
  toggleMute: () => void;
  children: ReactNode;
  setMemoOpen: (open: boolean) => void;
  setAssistantOpen: (open: boolean) => void;
  onPriorityHighClick: () => void;
};

export const SessionBottomNavigationTemplate = ({ value, isMute, toggleMute, children, setMemoOpen, setAssistantOpen, onPriorityHighClick }: sessionBottomNavigationTemplateProps) => {
  return (
    <Box width={"100vw"} height={"100vh"}>
      <Box height={"calc(100vh - 70px)"} overflow={"auto"}>
        {children}
      </Box>
      <SessionBottomNavigation value={value} isMute={isMute} toggleMute={toggleMute} setMemoOpen={setMemoOpen} setAssistantOpen={setAssistantOpen} onPriorityHighClick={onPriorityHighClick} />
    </Box>
  );
};
