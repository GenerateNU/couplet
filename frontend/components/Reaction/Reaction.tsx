import React from "react";
import { View } from "react-native";
import ReactionButton from "./ReactionButton";
import { ReactionProps } from "./ReactionProps";

function Reaction({ handleReact }: ReactionProps) {
  return (
    <View
      style={{
        flexDirection: "row",
        marginHorizontal: "5%",
        justifyContent: "space-between",
        paddingTop: "5%"
      }}
    >
      <ReactionButton like={false} icon="cancel" label="Nope" handleReact={handleReact} />
      <ReactionButton like icon="heart" label="Save" handleReact={handleReact} />
    </View>
  );
}

export default Reaction;
