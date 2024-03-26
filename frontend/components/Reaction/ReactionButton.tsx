import React from "react";
import { IconButton } from "react-native-paper";
import COLORS from "../../colors";

/* eslint-disable react/no-children-prop */

export type ReactionButtonProps = {
  like: boolean;
  icon: string;
  handleReact: (like: boolean) => void;
};

function ReactionButton({ like, icon, handleReact }: ReactionButtonProps) {
  return (
    <IconButton
      icon={icon}
      containerColor={like ? COLORS.primary : "#f3f6f4"}
      iconColor={like ? "white" : "black"}
      mode="contained"
      size={36}
      accessibilityLabel={like ? "Not interested" : "Like"}
      onPress={() => handleReact(like)}
    />
  );
}

export default ReactionButton;
