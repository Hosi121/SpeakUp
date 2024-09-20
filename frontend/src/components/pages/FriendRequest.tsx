import { useState, useEffect } from 'react';
import { Box, Typography, Avatar, Button, Paper, List, ListItem, ListItemAvatar, ListItemText, Container, IconButton } from "@mui/material";
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

// カスタムフック: フレンド情報の取得
const useFetchFriends = (userIds: number[]) => {
  const [friends, setFriends] = useState<Friend[]>([]);

  useEffect(() => {
    const fetchFriends = async () => {
      try {
        const friendsData = await Promise.all(
          userIds.map(async (id) => {
            const response = await api.get(`/users/${id}/avatar`);
            return {
              id,
              username: `User ${id}`, // 実際のユーザー名を取得するAPIがある場合は、そちらを使用
              avatar: response.data.avatarURL,
              isFriend: false, // 初期状態では全てfalse
            };
          })
        );
        setFriends(friendsData);
      } catch (error) {
        console.error('Failed to fetch friends:', error);
      }
    };

    fetchFriends();
  }, [userIds]);

  return friends;
};

// カスタムフック: フレンドリクエストの送信
const useSendFriendRequest = () => {
  const sendRequest = async (targetUserId: number) => {
    try {
      await api.post('/friend/register', {
        target_user_id: targetUserId,
      });
      return true;
    } catch (error) {
      console.error('Failed to send friend request:', error);
      return false;
    }
  };

  return sendRequest;
};

const FriendRequestComponent: React.FC = () => {
  const navigate = useNavigate();
  const userIds = [2, 3]; // セッションしたユーザーのID
  const friends = useFetchFriends(userIds);
  const sendFriendRequest = useSendFriendRequest();

  const handleEndSession = () => {
    navigate("/home");
  };

  const handleGoBack = () => {
    navigate(-1);
  };

  const handleFriendRequest = async (friendId: number) => {
    const success = await sendFriendRequest(friendId);
    if (success) {
      // 成功時の処理（例：状態の更新）
      console.log('Friend request sent successfully');
      // フレンドリストを更新する（該当するフレンドのisFriendをtrueに設定）
      setFriends(prevFriends =>
        prevFriends.map(friend =>
          friend.id === friendId ? { ...friend, isFriend: true } : friend
        )
      );
    }
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
    </Container>
  );
};

export const FriendRequest: React.FC = () => {
  return <FriendRequestComponent />;
};

export default FriendRequest;