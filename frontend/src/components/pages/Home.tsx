import { Stack } from "@mui/system";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";

const HomeContainer = () => {
  return <Stack>home</Stack>;
};

export const Home = () => {
  return (
    <BottomNavigationTemplate value="home">
      <HomeContainer />
    </BottomNavigationTemplate>
  );
};
