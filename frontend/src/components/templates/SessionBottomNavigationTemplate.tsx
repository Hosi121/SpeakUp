import { ReactNode } from "react";
import { SessionBottomNavigation, sessionBottomNavigation } from "../utils/SessionBottomNavigation";
import { Box } from "@mui/material";

type sessionBottomNavigationTemplateProps = {
  value: sessionBottomNavigation;
  isMute: boolean;
  children: ReactNode;
  setMemoOpen: (open: boolean) => void;
  setAssistantOpen: (open: boolean) => void;
};

export const SessionBottomNavigationTemplate = ({ value, isMute, children, setMemoOpen, setAssistantOpen }: sessionBottomNavigationTemplateProps) => {
  return (
    <Box width={"100vw"} height={"100vh"}>
      <Box height={"calc(100vh - 70px)"} overflow={"auto"}>
        {children}
      </Box>
      <SessionBottomNavigation value={value} isMute={isMute} setMemoOpen={setMemoOpen} setAssistantOpen={setAssistantOpen} />
    </Box>
  );
};
