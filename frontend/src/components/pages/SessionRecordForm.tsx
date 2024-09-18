import React, { useState } from 'react';
import { Box, Typography, TextField, Button, Paper } from '@mui/material';
import { styled } from '@mui/system';
import { useNavigate } from 'react-router-dom';

const YellowBox = styled(Box)({
  backgroundColor: '#FFD700',
  padding: '10px',
  marginBottom: '20px',
});

const DashedBox = styled(Paper)({
  border: '2px dashed #FFD700',
  padding: '20px',
  marginBottom: '20px',
});

const NextSessionBox = styled(Box)({
  backgroundColor: '#FFD700',
  padding: '10px',
  marginBottom: '20px',
});

const PinkButton = styled(Button)({
  backgroundColor: '#FF69B4',
  color: 'white',
  '&:hover': {
    backgroundColor: '#FF1493',
  },
});

const SessionRecordForm: React.FC = () => {
  const [satisfaction, setSatisfaction] = useState<string>('');
  const [thoughts, setThoughts] = useState<string>('');
  const [learnedExpressions, setLearnedExpressions] = useState<string>('');
  const navigate = useNavigate();

  const handleNext = () => {
    navigate('/friendrequest');
  };

  return (
    <Box sx={{ maxWidth: 400, margin: 'auto', padding: 2 }}>
      <YellowBox>
        <Typography variant="h4" align="center" sx={{ fontWeight: 'bold' }}>
          GOOD JOB!
        </Typography>
      </YellowBox>
      
      <Typography variant="h6" align="center" gutterBottom>
        記録を残してみませんか？
      </Typography>
      
      <DashedBox elevation={0}>
        <Typography variant="body1">5月15日</Typography>
        <Typography variant="body1">参加したセッション数: 3</Typography>
        <TextField
          fullWidth
          label="満足度 (%)"
          value={satisfaction}
          onChange={(e) => setSatisfaction(e.target.value)}
          margin="normal"
        />
        <TextField
          fullWidth
          label="感想"
          multiline
          rows={3}
          value={thoughts}
          onChange={(e) => setThoughts(e.target.value)}
          margin="normal"
        />
        <TextField
          fullWidth
          label="学んだ表現"
          multiline
          rows={3}
          value={learnedExpressions}
          onChange={(e) => setLearnedExpressions(e.target.value)}
          margin="normal"
        />
      </DashedBox>
      
      <Typography variant="h6" align="center" gutterBottom>
        次に参加予定のセッション
      </Typography>
      
      <NextSessionBox>
        <Typography variant="body1" align="center">
          7月17日20:00〜
        </Typography>
      </NextSessionBox>
      
      <PinkButton variant="contained" fullWidth onClick={handleNext}>
        次へ
      </PinkButton>
    </Box>
  );
};

export default SessionRecordForm;