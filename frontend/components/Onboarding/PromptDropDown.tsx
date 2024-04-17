import React, { useEffect, useState } from "react";
import { StyleSheet } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";
import { screenHeight } from "../../utils/dimensions";

interface DropDownLocationProps {
  onPromptChange: (local: string) => void;
  selectedPrompt: string;
}

function PromptDropDown({ onPromptChange, selectedPrompt }: DropDownLocationProps) {
  const [open, setOpen] = useState(false);
  const [prompt, setPrompt] = useState(selectedPrompt);
  const [items, setItems] = useState([
    { label: "My ideal date...", value: "My ideal date..." },
    { label: "My perfect day consists of...", value: "My perfect day consists of..." },
    { label: "On weekends you can find me...", value: "On weekends you can find me..." },
    { label: "What I'm looking for on this app...", value: "What I'm looking for on this app..." },
    { label: "A story you should know about me...", value: "A story you should know about me..." }
  ]);

  useEffect(() => {
    onPromptChange(prompt);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [prompt]);

  return (
    <DropDownPicker
      open={open}
      value={prompt}
      items={items}
      style={scaledStyles.promptBox}
      setOpen={setOpen}
      setValue={setPrompt}
      setItems={setItems}
      dropDownContainerStyle={{ height: screenHeight * 0.15 }}
    />
  );
}

const styles = StyleSheet.create({
  charCount: {
    textAlign: "right"
  },
  promptBox: {
    marginBottom: 10,
    borderRadius: 10
  },
  responseBox: {
    padding: 10,
    borderWidth: 1,
    borderColor: COLORS.darkGray,
    borderRadius: 10,
    height: "40%"
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default PromptDropDown;
