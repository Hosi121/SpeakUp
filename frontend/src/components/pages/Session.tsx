import { Card, CardContent, Typography, Avatar, Grid, Box, Container, Stack } from "@mui/material";
import { Favorite, Person } from "@mui/icons-material";
import TopSection from "../utils/TopSection";

function Session() {
  const users = [
    { name: "User1", icon: <Person />, description: "英語" },
    { name: "User2", icon: <Favorite />, description: "苗字" },
  ];

  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "100vh",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <TopSection />
        <Stack sx={{ margin: "30px auto 0", width: "100%" }}>
          <Typography variant="h5" align="center" gutterBottom color="primary.main" fontWeight="bolder">
            テーマ: 好きな言葉
          </Typography>
          <Box sx={{ display: "flex", flexDirection: "column", justifyContent: "center", mb: 2 }}>
            {users.map((user, index) => (
              <Card key={index} sx={{ bgcolor: "secondary.main", mb: 2, width: "100%", height: "25vh", borderRadius: 5, border: "6px solid #eee", boxSizing: "border-box", p: 1, display: "grid", placeContent: "center" }}>
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
}

export default Session;
