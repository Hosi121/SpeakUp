import { Card, CardContent, Typography, Avatar, Grid, Box } from '@mui/material';
import { Favorite, Person } from '@mui/icons-material';

function Session() {
  const users = [
    { name: "User1", icon: <Person />, description: "英語" },
    { name: "User2", icon: <Favorite />, description: "苗字" }
  ];

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', p: 2 }}>
      <Typography variant="h4" align="center" gutterBottom>
        好きな言葉
      </Typography>
      {users.map((user, index) => (
        <Card key={index} sx={{ width: 300, bgcolor: '#fff59d', mb: 2 }}>
          <CardContent>
            {/* カード内のテーマを削除しました */}
            <Grid container spacing={2} justifyContent="center">
              <Grid item>
                <Avatar sx={{ bgcolor: 'transparent', width: 56, height: 56 }}>
                  {user.icon}
                </Avatar>
                <Typography variant="subtitle1" align="center">{user.name}</Typography>
              </Grid>
              <Grid item xs={12}>
                <Typography variant="body2" align="center">{user.description}</Typography>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
      ))}
    </Box>
  );
}

export default Session;