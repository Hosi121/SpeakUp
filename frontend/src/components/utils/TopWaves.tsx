import styles from "./TopWaves.module.css";
import { Box } from "@mui/material";

const TopWaves = ({ isFlipped }: { isFlipped: boolean }) => {
  return (
    <Box sx={{ margin: "0 calc(50% - 50vw)", width: "100vw" }}>
      <svg className={`${styles.waves} ${isFlipped ? styles.flipped : ""}`} xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" viewBox="0 24 150 28" preserveAspectRatio="none" shape-rendering="auto">
        <defs>
          <path id="gentle-wave" d="M-160 44c30 0 58-18 88-18s 58 18 88 18 58-18 88-18 58 18 88 18 v44h-352z" />
        </defs>
        <g className={styles.parallax}>
          <use xlinkHref="#gentle-wave" x="48" y="0" fill="rgba(255,255,255,0.7)" />
          <use xlinkHref="#gentle-wave" x="48" y="3" fill="rgba(255,255,255,0.5)" />
          <use xlinkHref="#gentle-wave" x="48" y="5" fill="rgba(255,255,255,0.3)" />
          <use xlinkHref="#gentle-wave" x="48" y="7" fill="#fff" />
        </g>
      </svg>
    </Box>
  );
};

export default TopWaves;
