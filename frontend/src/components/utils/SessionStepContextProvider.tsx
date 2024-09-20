import { createContext, Dispatch, ReactNode, SetStateAction, useState } from 'react';

type SessionStepContextType = {
    sessionStep: number;
    setSessionStep: Dispatch<SetStateAction<number>>;
    learnedExpressions: string;
    setLearnedExpressions: Dispatch<SetStateAction<string>>;
}

export const SessionStepContext = createContext<SessionStepContextType>({} as SessionStepContextType);

export const SessionStepContextProvider = (props: { children: ReactNode }) => {
    const [sessionStep, setSessionStep] = useState(0);
    const [learnedExpressions, setLearnedExpressions] = useState("");
    return (
        <SessionStepContext.Provider value={{ sessionStep, setSessionStep, learnedExpressions, setLearnedExpressions }}>
            {props.children}
        </SessionStepContext.Provider >
    )
}

