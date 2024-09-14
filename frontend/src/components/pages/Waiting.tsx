import { Box, Button, Container, Typography } from '@mui/material';
import CircularProgress from '@mui/material/CircularProgress';
import Memo from './Memo';

const Waiting = () => {
  return (
    <Container sx={{ height: '100vh', display: 'flex', flexDirection: 'column', justifyContent: 'center', alignItems: 'center', padding: 3 }}>
      <Typography variant="h6" sx={{ mb: 2, backgroundColor: '#FFD700', borderRadius: 1, padding: 1, textAlign: 'center' }}>
        6月30日21:00〜
      </Typography>

      <Box sx={{ width: '100%', maxWidth: 500, bgcolor: 'background.paper', borderRadius: 1, p: 2, boxShadow: 3 }}>
        <Typography variant="subtitle1" gutterBottom component="div">
          持ち込みメモ
        </Typography>
        <Memo />
      </Box>

      <Box sx={{ mt: 2, display: 'flex', justifyContent: 'center', width: '100%' }}>
        <CircularProgress size={24} />
        <Typography variant="caption" sx={{ ml: 1 }}>
          Now loading...
        </Typography>
      </Box>

      <Button variant="contained" sx={{ mt: 4, bgcolor: 'secondary.main', '&:hover': { bgcolor: 'secondary.dark' } }}>
        セッション開始まで<br/>もうしばらくお待ちください
      </Button>
    </Container>
  );
};

export default Waiting;