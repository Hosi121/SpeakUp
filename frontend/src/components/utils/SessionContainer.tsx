import { Card, CardContent, Typography, Avatar, Grid, Box, Container, Stack } from "@mui/material";
import SessionCountDownModal from "./SessionCountDownModal";

const SessionContainer = ({ theme, users }: { theme: string; users: { name: string; icon: JSX.Element; description: string }[] }) => {
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
            {users.map((user, index) => (
              <Card
                key={index}
                sx={{
                  bgcolor: "secondary.main",
                  mb: 2,
                  width: "100%",
                  height: "25vh",
                  borderRadius: 5,
                  border: "6px solid #eee",
                  boxSizing: "border-box",
                  p: 1,
                  display: "grid",
                  placeContent: "center",
                }}
              >
                <CardContent>
                  <Grid container spacing={2} justifyContent="center">
                    <Grid item>
                      <Avatar sx={{ bgcolor: "#eee", width: "80px", height: "80px" }}>{user.icon}</Avatar>
                      <Typography variant="h6" align="center" sx={{ mt: 1 }}>
                        {user.name}
                      </Typography>
                    </Grid>
                  </Grid>
                </CardContent>
              </Card>
            ))}
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};

export default SessionContainer;
