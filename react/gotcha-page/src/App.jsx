import logo from './logo.svg';
import './App.css';
import { useState } from 'react';

function App() {
  const [nicknameInputVal, setNicknameInputVal] = useState("")
  const onInput = (e) => {
    setNicknameInputVal(e.target.value)
  }
  return (
      <>
        <input value={nicknameInputVal} onChange={onInput}/>
      </>
  );
}

export default App;
