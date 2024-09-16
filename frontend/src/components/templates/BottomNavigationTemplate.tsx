import { ReactNode } from "react";
import {
  MainBottomNavigation,
  mainBottomNavigation,
} from "../utils/MainBottomNavigation";
import styles from "./BottomNavigationTemplate.module.css";

type bottomNavigationTemplateProps = {
  value: mainBottomNavigation;
  children: ReactNode;
};

export const BottomNavigationTemplate = ({
  value,
  children,
}: bottomNavigationTemplateProps) => {
  return (
    <div className={styles.container}>
      {children}
      <MainBottomNavigation value={value} />
    </div>
  );
};
