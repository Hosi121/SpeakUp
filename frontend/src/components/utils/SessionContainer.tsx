import { Card, CardContent, Typography, Avatar, Box, Container, Stack } from "@mui/material";
import Grid from '@mui/material/Grid2';
import SessionCountDownModal from "./SessionCountDownModal";

//      <SessionContainer theme={theme} users={users} isSpeak={isSpeak} isOpponentSpeak={isOpponentSpeak} />
type SessionContainerProps = {
  theme: string;
  users: { name: string; icon: JSX.Element; description: string }[]
  isSpeak: boolean;
  isOpponentSpeak: boolean;
}

const SessionContainer = ({ theme, users, isSpeak, isOpponentSpeak }: SessionContainerProps) => {
  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "100%",
        position: "relative",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <Stack sx={{ margin: "30px auto 0", width: "100%" }}>
          <Box
            sx={{
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              zIndex: 1000,
            }}
          >
            <SessionCountDownModal />
          </Box>
          <Typography variant="h5" align="center" gutterBottom color="primary.main" fontWeight="bolder">
            テーマ: {theme}
          </Typography>
          <Box sx={{ display: "flex", flexDirection: "column", justifyContent: "center", mb: 2 }}>
            <UserCard user={users[0]} isSpeak={isSpeak} index={0} />
            <UserCard user={users[1]} isSpeak={isOpponentSpeak} index={1} />
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};

export default SessionContainer;

type UserCardProps = {
  user: { name: string; icon: JSX.Element; description: string };
  isSpeak: boolean;
  index: number;
};

const UserCard = ({ user, isSpeak, index }: UserCardProps) => {
  return (
    <Card
      key={index}
      sx={{
        bgcolor: "secondary.main",
        mb: 2,
        width: "100%",
        height: "25vh",
        borderRadius: 5,
        border: "6px solid",
        boxSizing: "border-box",
        p: 1,
        display: "grid",
        placeContent: "center",
        borderColor: isSpeak ? "primary.main" : "#eee",
      }}
    >
      <CardContent>
        <Grid container spacing={2} justifyContent="center">
          <Grid>
            <Avatar sx={{ bgcolor: "#eee", width: "80px", height: "80px" }}>{user.icon}</Avatar>
            <Typography variant="h6" align="center" sx={{ mt: 1 }}>
              {user.name}
            </Typography>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  )
};
