import React from 'react';
import { useNavigate } from 'react-router-dom';
import { 
  Container, 
  Typography, 
  Radio, 
  RadioGroup, 
  FormControlLabel, 
  FormControl, 
  FormLabel, 
  Button, 
  Box 
} from '@mui/material';
import TopSection from '../utils/TopSection';

const SessionFeedback: React.FC = () => {
  const navigate = useNavigate();

  const handleNext = () => {
    navigate('/home');
  };

  return (
    <Container maxWidth="sm" sx={{ pt: 2, pb: 2 }}>
      <TopSection />
      <Typography variant="h4" component="h1" sx={{ color: '#FFD700', fontWeight: 'bold', my: 2 }}>
        EXCELLENT!
      </Typography>
      <Typography variant="h5" component="h2" sx={{ mb: 2 }}>
        最後に振り返りをしよう！
      </Typography>

      <Box sx={{ bgcolor: '#FFF9C4', p: 2, borderRadius: 2 }}>
        <FormControl component="fieldset">
          <FormLabel component="legend">楽しかった？</FormLabel>
          <RadioGroup defaultValue="楽しかった">
            <FormControlLabel value="すごく楽しかった" control={<Radio />} label="すごく楽しかった" />
            <FormControlLabel value="楽しかった" control={<Radio />} label="楽しかった" />
            <FormControlLabel value="あまり楽しくなかった" control={<Radio />} label="あまり楽しくなかった" />
          </RadioGroup>
        </FormControl>

        <FormControl component="fieldset" sx={{ mt: 2 }}>
          <FormLabel component="legend">英語だけで話せた？</FormLabel>
          <RadioGroup defaultValue="ほぼ英語だった">
            <FormControlLabel value="全部英語だった" control={<Radio />} label="全部英語だった" />
            <FormControlLabel value="ほぼ英語だった" control={<Radio />} label="ほぼ英語だった" />
            <FormControlLabel value="ときどき日本語だった" control={<Radio />} label="ときどき日本語だった" />
          </RadioGroup>
        </FormControl>

        <FormControl component="fieldset" sx={{ mt: 2 }}>
          <FormLabel component="legend">前回よりたくさん話せた？</FormLabel>
          <RadioGroup defaultValue="同じくらい">
            <FormControlLabel value="前回より多く話せた" control={<Radio />} label="前回より多く話せた" />
            <FormControlLabel value="同じくらい" control={<Radio />} label="同じくらい" />
            <FormControlLabel value="前回より少なかった" control={<Radio />} label="前回より少なかった" />
          </RadioGroup>
        </FormControl>

        <FormControl component="fieldset" sx={{ mt: 2 }}>
          <FormLabel component="legend">相手は英語だけで話してた？</FormLabel>
          <RadioGroup defaultValue="ほとんど英語だった">
            <FormControlLabel value="全部英語だった" control={<Radio />} label="全部英語だった" />
            <FormControlLabel value="ほとんど英語だった" control={<Radio />} label="ほとんど英語だった" />
            <FormControlLabel value="ときどき日本語だった" control={<Radio />} label="ときどき日本語だった" />
          </RadioGroup>
        </FormControl>

        <FormControl component="fieldset" sx={{ mt: 2 }}>
          <FormLabel component="legend">相手はたくさん話してくれた？</FormLabel>
          <RadioGroup defaultValue="まあまあ">
            <FormControlLabel value="はい" control={<Radio />} label="はい" />
            <FormControlLabel value="まあまあ" control={<Radio />} label="まあまあ" />
            <FormControlLabel value="いいえ" control={<Radio />} label="いいえ" />
          </RadioGroup>
        </FormControl>

        <FormControl component="fieldset" sx={{ mt: 2 }}>
          <FormLabel component="legend">相手はたくさん聞いてくれた？</FormLabel>
          <RadioGroup defaultValue="まあまあ">
            <FormControlLabel value="はい" control={<Radio />} label="はい" />
            <FormControlLabel value="まあまあ" control={<Radio />} label="まあまあ" />
            <FormControlLabel value="いいえ" control={<Radio />} label="いいえ" />
          </RadioGroup>
        </FormControl>
      </Box>

      <Typography variant="h6" sx={{ mt: 3, mb: 1 }}>
        次に参加予定のセッション
      </Typography>
      <Box sx={{ bgcolor: '#FFD700', p: 2, borderRadius: 2, mb: 3 }}>
        <Typography variant="h5" sx={{ fontWeight: 'bold' }}>
          7月17日20:00〜
        </Typography>
      </Box>

      <Button 
        variant="contained" 
        color="primary" 
        fullWidth 
        onClick={handleNext}
        sx={{ mt: 2 }}
      >
        次へ
      </Button>
    </Container>
  );
};

export default SessionFeedback;