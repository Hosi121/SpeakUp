import { ReactNode } from "react";
import { SessionBottomNavigation, sessionBottomNavigation } from "../utils/SessionBottomNavigation";
import { Box } from "@mui/material";

type sessionBottomNavigationTemplateProps = {
  value: sessionBottomNavigation;
  isMute: boolean;
  children: ReactNode;
};

export const SessionBottomNavigationTemplate = ({ value, isMute, children }: sessionBottomNavigationTemplateProps) => {
  return (
    <Box width={"100vw"} height={"100vh"}>
      <Box height={"calc(100vh - 70px)"} overflow={"auto"}>
        {children}
      </Box>
      <SessionBottomNavigation value={value} isMute={isMute} />
    </Box>
  );
};
