import { useEffect, useState } from "react";
import { Avatar, Button, List, ListItem, ListItemAvatar, ListItemText } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { fetchFriendList } from "../../services/friendService";
import type { FriendSummary } from "../../types/types";

const FriendList: React.FC = () => {
  const [friends, setFriends] = useState<FriendSummary[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    // フレンドリストを取得
    const loadFriends = async () => {
      try {
        const data = await fetchFriendList();
        setFriends(data);
      } catch (error) {
        console.error("Failed to fetch friend list", error);
      }
    };
    loadFriends();
  }, []);

  const handleMessage = (friendName: string) => {
    // メッセージページに遷移
    navigate(`/message/${friendName}`);
  };

  return (
    <List>
      {friends.map((friend) => (
        <ListItem key={friend.id} sx={{ mb: 3, width: "100%", p: 0 }}>
          <ListItemAvatar sx={{ mr: 1 }}>
            <Avatar src={friend.avatarUrl} alt={friend.username} sx={{ width: "50px", height: "50px" }} />
          </ListItemAvatar>
          <ListItemText primary={friend.username} primaryTypographyProps={{ fontSize: "1.5rem" }} />
          <Button variant="contained" onClick={() => handleMessage(friend.username)} sx={{ borderRadius: 3 }}>
            メッセージ
          </Button>
        </ListItem>
      ))}
    </List>
  );
};

export default FriendList;
