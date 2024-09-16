import { Stack } from "@mui/system";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";

const SessionListContainer = () => {
  return <Stack>sessionList</Stack>;
};

export const SessionList = () => {
  return (
    <BottomNavigationTemplate value="session">
      <SessionListContainer />
    </BottomNavigationTemplate>
  );
};
