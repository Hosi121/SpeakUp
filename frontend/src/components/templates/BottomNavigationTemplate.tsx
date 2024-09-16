import { ReactNode } from "react";
import { MainBottomNavigation, mainBottomNavigation } from "../utils/MainBottomNavigation";
import { Box } from "@mui/material";

type bottomNavigationTemplateProps = {
  value: mainBottomNavigation;
  children: ReactNode;
};

export const BottomNavigationTemplate = ({ value, children }: bottomNavigationTemplateProps) => {
  return (
    <Box width={"100vw"} height={"calc(100vh - 70px )"}>
      {children}
      <MainBottomNavigation value={value} />
    </Box>
  );
};
