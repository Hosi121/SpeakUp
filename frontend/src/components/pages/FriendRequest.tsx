import { useState, useEffect, useCallback } from 'react';
import { Box, Typography, Avatar, Button, Paper, List, ListItem, ListItemAvatar, ListItemText, Container, IconButton, Snackbar, CircularProgress } from "@mui/material";
import { ArrowBack } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";
import TopSection from "../utils/TopSection";
import api from "../../services/api";

interface Friend {
  id: number;
  username: string;
  avatar: string;
  isFriend: boolean;
}

const FriendRequestComponent: React.FC = () => {
  const navigate = useNavigate();
  const [friends, setFriends] = useState<Friend[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
  const userIds = [2, 3]; // セッションしたユーザーのID

  const fetchUserInfo = async (userId: number): Promise<Friend> => {
    try {
      const [userResponse, avatarResponse] = await Promise.all([
        api.get(`/users/search/id/${userId}`),
        api.get(`/users/${userId}/avatar`)
      ]);
      
      return {
        id: userId,
        username: userResponse.data.username,
        avatar: avatarResponse.data.avatarURL,
        isFriend: false,
      };
    } catch (error: any) {
      console.error(`Failed to fetch user info for user ${userId}:`, error);
      throw error;
    }
  };

  const fetchFriends = useCallback(async () => {
    setLoading(true);
    try {
      const friendsData = await Promise.all(userIds.map(fetchUserInfo));
      setFriends(friendsData);
    } catch (error: any) {
      console.error('Failed to fetch friends:', error);
      setError(`Failed to fetch friends: ${error.message}`);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchFriends();
  }, [fetchFriends]);

  const sendFriendRequest = async (targetUserId: number) => {
    try {
      const response = await api.post('/friend/register', {
        target_user_id: targetUserId,
      });
      console.log('Friend request response:', response.data);
      return true;
    } catch (error: any) {
      console.error('Failed to send friend request:', error.response?.data || error.message);
      setError(`Failed to send friend request: ${error.response?.data?.error || error.message}`);
      return false;
    }
  };

  const handleEndSession = () => {
    navigate("/home");
  };

  const handleGoBack = () => {
    navigate(-1);
  };

  const handleFriendRequest = async (friendId: number) => {
    const success = await sendFriendRequest(friendId);
    if (success) {
      console.log('Friend request sent successfully');
      setFriends(prevFriends =>
        prevFriends.map(friend =>
          friend.id === friendId ? { ...friend, isFriend: true } : friend
        )
      );
    }
  };

  const handleCloseError = () => {
    setError(null);
  };

  if (loading) {
    return (
      <Container sx={{ display: "flex", justifyContent: "center", alignItems: "center", height: "100vh" }}>
        <CircularProgress />
      </Container>
    );
  }

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
          <Typography variant="h6" sx={{ ml: 2 }} fontWeight="bolder">
            フレンド申請
          </Typography>
        </Box>

        <Typography variant="body2" sx={{ marginBottom: 2 }} textAlign="left">
          今日話した相手ともっと話したいときは、 フレンド申請をしてメッセージでセッション の続きを話そう！
        </Typography>

        <List sx={{ maxHeight: "50vh", overflow: "auto" }}>
          {friends.map((friend) => (
            <Paper key={friend.id} elevation={3} sx={{ marginBottom: 2, padding: 2 }}>
              <ListItem alignItems="flex-start" sx={{ padding: 0 }}>
                <ListItemAvatar>
                  <Avatar src={friend.avatar} alt={friend.username} />
                </ListItemAvatar>
                <ListItemText
                  primary={friend.username}
                  secondary={
                    <>
                      <Typography component="span" variant="body2" color="text.primary">
                        User ID: {friend.id}
                      </Typography>
                    </>
                  }
                />
              </ListItem>
              <Box sx={{ display: "flex", justifyContent: "space-between", marginTop: 2 }}>
                <Button 
                  variant="contained" 
                  disabled={friend.isFriend}
                  onClick={() => handleFriendRequest(friend.id)}
                >
                  {friend.isFriend ? "フレンド申請済" : "フレンド申請"}
                </Button>
                <Button variant="contained">メッセージ</Button>
              </Box>
            </Paper>
          ))}
        </List>
      </Container>

      <Box sx={{ p: 2 }}>
        <Typography variant="body2" sx={{ marginBottom: 2 }} textAlign="left">
          フレンド申請は、 記録＞セッション履歴からでもできるよ！
        </Typography>

        <Button variant="contained" fullWidth onClick={handleEndSession}>
          セッションを終わる→
        </Button>
      </Box>

      <Snackbar
        open={!!error}
        autoHideDuration={6000}
        onClose={handleCloseError}
        message={error}
      />
    </Container>
  );
};

export const FriendRequest: React.FC = () => {
  return <FriendRequestComponent />;
};

export default FriendRequest;