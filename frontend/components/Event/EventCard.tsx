import React, { useEffect, useState } from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { Button, Icon } from "react-native-paper";
import { getEvents } from "../../api/events";
import { getOrgById } from "../../api/orgs";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";
import OrgTag from "./OrgTag";

type Event = Awaited<ReturnType<typeof getEvents>>[number];
type Org = Awaited<ReturnType<typeof getOrgById>>;

const IMAGE = require("../../assets/profile.png");

export type EventCardProps = {
  event: Event;
  handleReact: (like: boolean) => void;
};

export default function EventCard({ handleReact, event }: EventCardProps) {
  const [org, setOrg] = useState<Org>();

  useEffect(() => {
    console.log("orgId", event);

    if (!event.orgId) return;
    getOrgById({ id: event.orgId })
      .then((fetchedOrg) => setOrg(fetchedOrg))
      .catch((e) => console.error(e));
  }, [event]);

  return (
    <View style={scaledStyles.container}>
      <View style={scaledStyles.detail}>
        <Icon source="map-marker" size={24} color={COLORS.darkPurple} />
        <Text style={{ fontSize: 18, marginRight: 24, fontFamily: "DMSansRegular" }}>
          Frog Pond
        </Text>
      </View>
      <View style={scaledStyles.detail}>
        <Icon source="calendar-blank" size={24} color={COLORS.darkPurple} />
        <Text style={{ fontSize: 18 }}>Open today · Closes at 10:00 PM</Text>
      </View>

      <View style={scaledStyles.detail}>
        <Icon source="currency-usd" size={24} color={COLORS.darkPurple} />
        <Text style={{ fontSize: 18, marginRight: 20, fontFamily: "DMSansRegular" }}>30</Text>
      </View>
      <Text style={scaledStyles.eventBio}>{event?.bio}</Text>
      <View style={scaledStyles.viewShare}>
        <Button
          mode="outlined"
          buttonColor={COLORS.white}
          textColor={COLORS.primary}
          labelStyle={scaledStyles.buttonLabel}
          style={{ borderColor: COLORS.primary, borderWidth: 2 }}
        >
          View details
        </Button>
        <Button
          mode="contained"
          icon="export-variant"
          buttonColor={COLORS.primary}
          textColor={COLORS.white}
          labelStyle={{ ...scaledStyles.buttonLabel, paddingHorizontal: 8, fontWeight: "700" }}
          contentStyle={{ flexDirection: "row-reverse" }}
        >
          Share event
        </Button>
      </View>
      <View style={scaledStyles.orgSection}>
        <View style={scaledStyles.orgNameSection}>
          <Image source={IMAGE} />
          <View>
            <Text style={scaledStyles.orgNameText}>{org?.name}</Text>
            <Text style={scaledStyles.orgHandleText}>@{org?.name?.replaceAll(" ", "")}</Text>
          </View>
        </View>
        <View style={scaledStyles.tags}>
          {org?.tags?.map((tag) => <OrgTag key={tag} text={tag} />)}
        </View>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexGrow: 1,
    paddingTop: 10,
    paddingBottom: 50
  },
  detail: {
    flexDirection: "row",
    columnGap: 8,
    marginBottom: 5
  },
  eventBio: { fontSize: 18, fontFamily: "DMSansRegular", marginVertical: 10 },
  orgSection: {
    justifyContent: "space-between",
    marginVertical: 10,
    paddingVertical: 20,
    borderColor: COLORS.lightGray,
    borderTopWidth: 1
  },
  orgNameSection: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: 15
  },
  orgNameText: { marginLeft: 15, fontSize: 18, fontFamily: "DMSansMedium" },
  orgHandleText: { marginLeft: 15, fontSize: 12, fontFamily: "DMSansRegular", fontWeight: "400" },
  tags: {
    flexDirection: "row",
    flexWrap: "wrap",
    alignContent: "stretch",
    columnGap: 15,
    rowGap: 10
  },
  viewShare: {
    flexDirection: "row",
    justifyContent: "space-between",
    marginVertical: 10
  },
  buttonLabel: {
    fontFamily: "DMSansMedium",
    fontSize: 16,
    paddingHorizontal: 16
  }
});

const scaledStyles = scaleStyleSheet(styles);
