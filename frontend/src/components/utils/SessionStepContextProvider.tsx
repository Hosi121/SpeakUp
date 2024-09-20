import { createContext, Dispatch, ReactNode, SetStateAction, useState } from 'react';

type SessionStepContextType = {
    sessionStep: number;
    setSessionStep: Dispatch<SetStateAction<number>>;
}

export const SessionStepContext = createContext<SessionStepContextType>({} as SessionStepContextType);

export const SessionStepContextProvider = (props: { children: ReactNode }) => {
    const [sessionStep, setSessionStep] = useState(0);
    return (
        <SessionStepContext.Provider value={{ sessionStep, setSessionStep }}>
            {props.children}
        </SessionStepContext.Provider >
    )
}

