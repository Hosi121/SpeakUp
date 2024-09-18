import React from 'react';
import { Box, Typography, Avatar, Button, Paper, List, ListItem, ListItemAvatar, ListItemText, Container, IconButton } from '@mui/material';
import { Person, Star, ArrowBack } from '@mui/icons-material';
import { styled } from '@mui/system';
import { useNavigate } from 'react-router-dom';
import TopSection from "../utils/TopSection";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";

interface Friend {
  id: number;
  username: string;
  avatar: string;
  topic: string;
  date: string;
  level: number;
  isFriend: boolean;
}

const YellowButton = styled(Button)({
  backgroundColor: '#FFD700',
  color: 'black',
  '&:hover': {
    backgroundColor: '#FFD700',
  },
});

const PinkButton = styled(Button)({
  backgroundColor: '#FF69B4',
  color: 'white',
  '&:hover': {
    backgroundColor: '#FF1493',
  },
});

const GrayButton = styled(Button)({
  backgroundColor: '#E0E0E0',
  color: 'black',
  '&:hover': {
    backgroundColor: '#E0E0E0',
  },
});

const FriendRequestComponent: React.FC = () => {
  const navigate = useNavigate();

  const friends: Friend[] = [
    {
      id: 1,
      username: 'Mike',
      avatar: '',
      topic: '好きなスポーツ',
      date: '2024/5/15',
      level: 3,
      isFriend: true,
    },
    {
      id: 2,
      username: 'ゆい',
      avatar: '',
      topic: '好きなスポーツ',
      date: '2024/5/15',
      level: 3,
      isFriend: false,
    },
  ];

  const handleEndSession = () => {
    navigate('/home');
  };

  const handleGoBack = () => {
    navigate(-1);
  };

  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "calc(100vh - 70px)",
        boxSizing: "border-box",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <TopSection />
        <Box
          sx={{
            position: "relative",
            display: "flex",
            justifyContent: "flex-start",
            alignItems: "center",
            pt: 2,
            pb: 2,
          }}
        >
          <IconButton onClick={handleGoBack}>
            <ArrowBack sx={{ fontSize: 40 }} />
          </IconButton>
          <Typography variant="h6" sx={{ ml: 2 }}>フレンド申請</Typography>
        </Box>

        <Typography variant="body2" sx={{ marginBottom: 2 }}>
          今日話した相手ともっと話したいときは、
          フレンド申請をしてメッセージでセッション
          の続きを話そう！
        </Typography>

        <List sx={{ maxHeight: "50vh", overflow: "auto" }}>
          {friends.map((friend) => (
            <Paper key={friend.id} elevation={3} sx={{ marginBottom: 2, padding: 2 }}>
              <ListItem alignItems="flex-start" sx={{ padding: 0 }}>
                <ListItemAvatar>
                  {friend.isFriend ? (
                    <Avatar sx={{ bgcolor: 'gold' }}>
                      <Star />
                    </Avatar>
                  ) : (
                    <Avatar src={friend.avatar} alt={friend.username} />
                  )}
                </ListItemAvatar>
                <ListItemText
                  primary={friend.username}
                  secondary={
                    <>
                      <Typography component="span" variant="body2" color="text.primary">
                        話したテーマ：{friend.topic}
                      </Typography>
                      <br />
                      <Typography component="span" variant="body2" color="text.primary">
                        話した日付：{friend.date}
                      </Typography>
                      <br />
                      <Typography component="span" variant="body2" color="text.primary">
                        レベル {friend.level}
                      </Typography>
                    </>
                  }
                />
              </ListItem>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', marginTop: 2 }}>
                <GrayButton variant="contained" disabled={friend.isFriend}>
                  {friend.isFriend ? 'フレンド申請済' : 'フレンド申請'}
                </GrayButton>
                <PinkButton variant="contained">メッセージ</PinkButton>
              </Box>
            </Paper>
          ))}
        </List>
      </Container>

      <Box sx={{ p: 2 }}>
        <Typography variant="body2" sx={{ marginBottom: 2 }}>
          フレンド申請は、
          記録＞セッション履歴からでもできるよ！
        </Typography>

        <Button variant="outlined" fullWidth onClick={handleEndSession}>
          セッションを終わる→
        </Button>
      </Box>
    </Container>
  );
};

export const FriendRequest: React.FC = () => {
  return (
    <FriendRequestComponent />
  );
};

export default FriendRequest;