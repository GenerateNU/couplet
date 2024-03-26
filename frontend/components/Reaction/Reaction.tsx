import React from "react";
import { View } from "react-native";
import ReactionButton from "./ReactionButton";

export type ReactionProps = {
  handleReact: (like: boolean) => void;
};

function Reaction({ handleReact }: ReactionProps) {
  return (
    <View
      style={{
        flexDirection: "row",
        justifyContent: "space-between",
        paddingHorizontal: 20
      }}
    >
      <ReactionButton like={false} icon="window-close" handleReact={handleReact} />
      <ReactionButton like icon="heart" handleReact={handleReact} />
    </View>
  );
}

export default Reaction;
