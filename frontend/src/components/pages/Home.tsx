import { Box, Container } from "@mui/material";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import TopSection from "../utils/TopSection";
import TopWaves from "../utils/TopWaves";
import HomeLogo from "../../assets/homeLogo";

const HomeContainer = () => {
  return (
    <Container sx={{ display: "flex", flexDirection: "column", justifyContent: "space-between", height: "calc(100vh - 70px)" }}>
      <Container sx={{ pt: 3 }}>
        <TopSection />
        <Container sx={{ mb: 3 }}>
          <TopWaves isFlipped={false} />
          <Box sx={{ margin: "0 calc(50% - 50vw)", width: "100vw", height: "15vh", backgroundColor: "secondary.main", display: "grid", placeItems: "center" }}>
            <HomeLogo style={{ width: "70%" }} />
          </Box>
          <TopWaves isFlipped={true} />
        </Container>
      </Container>
    </Container>
  );
};

export const Home = () => {
  return (
    <BottomNavigationTemplate value="home">
      <HomeContainer />
    </BottomNavigationTemplate>
  );
};
