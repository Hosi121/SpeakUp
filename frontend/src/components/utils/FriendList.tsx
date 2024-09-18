import { useEffect, useState } from "react";
import { Avatar, Button, List, ListItem, ListItemAvatar, ListItemText } from "@mui/material";
import { useNavigate } from "react-router-dom";
import api from "../../services/api";

interface Friend {
  id: number;
  username: string;
  avatar: string;
}

const FriendList: React.FC = () => {
  const [friends, setFriends] = useState<Friend[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    // フレンドリストを取得
    api
      .get<{ friends: Friend[] }>("/friend/me")
      .then((response) => {
        setFriends(response.data.friends);
      })
      .catch((error) => {
        console.error("Failed to fetch friend list", error);
      });
  }, []);

  const handleMessage = (friendName: string) => {
    // メッセージページに遷移
    navigate(`/message/${friendName}`);
  };

  console.log(friends);
  return (
    <List>
      {friends.map((friend) => (
        <ListItem key={friend.id} sx={{ mb: 3, width: "100%", p: 0 }}>
          <ListItemAvatar sx={{ mr: 1 }}>
            <Avatar src={friend.avatar} alt={friend.username} sx={{ width: "50px", height: "50px" }} />
          </ListItemAvatar>
          <ListItemText primary={friend.username} primaryTypographyProps={{ fontSize: "1.5rem" }} />
          <Button variant="contained" onClick={() => handleMessage(friend.username)}>
            メッセージ
          </Button>
        </ListItem>
      ))}
    </List>
  );
};

export default FriendList;
