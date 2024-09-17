import { useEffect, useState } from "react";
import {
  Avatar,
  Button,
  Container,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Typography,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import api from "../../services/api";

const FriendList = () => {
  const [friends, setFriends] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    // フレンドリストを取得
    api
      .get("/friend/me")
      .then((response) => {
        setFriends(response.data.friends);
      })
      .catch((error) => {
        console.error("Failed to fetch friend list", error);
      });
  }, []);

  const handleMessage = (friendName) => {
    // メッセージページに遷移
    navigate(`/message/${friendName}`);
  };

  return (
    <Container sx={{ pt: 3 }}>
      <Typography variant="h4" sx={{ mb: 4, fontWeight: "bold" }}>
        フレンド一覧
      </Typography>
      <List>
        {friends.map((friend) => (
          <ListItem key={friend.id} sx={{ mb: 2 }}>
            <ListItemAvatar>
              <Avatar src={friend.avatar} alt={friend.username} />
            </ListItemAvatar>
            <ListItemText primary={friend.username} />
            <Button
              variant="contained"
              onClick={() => handleMessage(friend.username)}
            >
              メッセージ
            </Button>
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default FriendList;