import React, { useEffect, useState } from "react";
import { ActivityIndicator, StyleSheet, Text, View } from "react-native";
import type { paths } from "../../api/schema";
import { getUserById } from "../../api/users";

const INSTAGRAM_ICON = require("../../assets/instagram.png");

const BLURRED_IG_USER = require("../../assets/blurredIgUser.png");

type User = paths["/users/{id}"]["get"]["responses"]["200"]["content"]["application/json"];

interface PersonProps {
  id: string;
  isMatched: boolean;
  likesYou: boolean;
  handleReact?: (like: boolean) => void;
}

export default function Person({ id, isMatched, likesYou, handleReact }: PersonProps) {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [profile, setProfile] = useState<User | null>(null);

  useEffect(() => {
    const load = async () => {
      setIsLoading(true);
      try {
        // This endpoint requires a LOT more information from the backend, including:
        // pronouns, location, school, work, height, bioQuestion, interests, relationshipType,
        // religion, politicalAffiliation, alcoholFrequency, smokingFrequency, drugFrequency, cannabisFrequency,
        // instagramUsername, mutualEvents, and captions
        const res = await getUserById({ id });
        setProfile(res);
      } catch (error) {
        throw new Error(String(error));
      }
      setIsLoading(false);
    };

    load();
  }, [id]);

  // const heightText = height ? `${height.feet}'${height.inches}"` : "";

  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size="large" />
      </View>
    );
  }

  const userNotFound = <Text>User not found :(</Text>;

  return userNotFound;

  // return profile ? <>
  //     <HeaderScrollView
  //       headerContainerStyle={{}}
  //       containerStyle={{
  //         flex: 1
  //       }}
  //       title={
  //         <>
  //           <Text style={{ fontFamily: "DMSansBold", marginRight: 10 }}>{profile.firstName} </Text>
  //           <Text id="hi" style={{ fontFamily: "DMSansRegular" }}>
  //             {profile.age}
  //           </Text>
  //         </>
  //       }
  //     >
  //       <Text
  //         style={{
  //           fontFamily: "DMSansRegular",
  //           color: "gray",
  //           fontSize: 14,
  //           marginLeft: 20,
  //           marginBottom: 10
  //         }}
  //       >
  //         {pronouns}
  //       </Text>

  //       <View>
  //         <View style={styles.infoContainer}>
  //           <Image style={styles.imageStyle} source={{ uri: profile.images[0] }} />
  //           <View style={styles.basicInfoContainer}>
  //             {location && (
  //               <View style={styles.basicInfoRow}>
  //                 <FontAwesomeIcon style={styles.basicInfoRowIcon} icon={faHouse} />
  //                 <Text style={styles.basicInfoRowText}>{location}</Text>
  //               </View>
  //             )}
  //             {school && (
  //               <View style={styles.basicInfoRow}>
  //                 <FontAwesomeIcon style={styles.basicInfoRowIcon} icon={faGraduationCap} />
  //                 <Text style={styles.basicInfoRowText}>{school}</Text>
  //               </View>
  //             )}
  //             {work && (
  //               <View style={styles.basicInfoRow}>
  //                 <FontAwesomeIcon style={styles.basicInfoRowIcon} icon={faBriefcase} />
  //                 <Text style={styles.basicInfoRowText}>{work}</Text>
  //               </View>
  //             )}
  //             {height && (
  //               <View style={styles.basicInfoRow}>
  //                 <FontAwesomeIcon style={styles.basicInfoRowIcon} icon={faRuler} />
  //                 <Text style={styles.basicInfoRowText}>{heightText}</Text>
  //               </View>
  //             )}
  //           </View>
  //           <View style={styles.promptContainer}>
  //             <Text style={{ fontFamily: "DMSansMedium", marginBottom: 5 }}>{promptQuestion}</Text>
  //             <Text style={styles.textStyle}>{profile.bio}</Text>
  //           </View>
  //           <View style={styles.separator} />
  //           <InfoChips items={interests} textColor="black" backgroundColor="lavender" />
  //           <Lifestyle
  //             relationshipType={relationshipType}
  //             religion={religion}
  //             politicalAffiliation={politicalAffiliation}
  //             alchoholFrequency={alchoholFrequency}
  //             smokingFrequency={smokingFrequency}
  //             drugFrequency={drugFrequency}
  //             cannabisFrequency={cannabisFrequency}
  //           />

  //           <View style={styles.separator} />
  //           <View
  //             style={{ flex: 1, justifyContent: "center", alignItems: "center", marginBottom: 20 }}
  //           >
  //             <Text style={styles.textStyle}>Match to unlock their Instagram!</Text>
  //             <View style={styles.instagramContainer}>
  //               <Image source={INSTAGRAM_ICON} style={{ width: 30, height: 30, marginRight: 10 }} />
  //               {isMatched && <Text style={{ fontFamily: "DMSansBold" }}>{instagramUsername}</Text>}
  //               {!isMatched && (
  //                 <Image source={BLURRED_IG_USER} style={{ width: "80%", height: 30 }} />
  //               )}
  //             </View>
  //           </View>
  //           <View>
  //             <Text style={styles.textStyle}>For our first date, let&apos;s go to...</Text>
  //             <ScrollView horizontal>
  //               {mutualEvents.map((event) => (
  //                 <View style={styles.mutualEventItemStyling}>
  //                   <EventCardItem
  //                     title={event.title}
  //                     description={event.description}
  //                     imageUrl={event.imageUrl}
  //                   />
  //                 </View>
  //               ))}
  //             </ScrollView>
  //           </View>
  //           <View style={styles.separator} />
  //           <View>
  //             {profile.images.slice(1).map((image) => (
  //               <View style={{ marginBottom: 20 }}>
  //                 <Image source={{ uri: image }} style={styles.imageStyle} />
  //                 <Text style={{ ...styles.textStyle, marginTop: 5 }}>{captions}</Text>
  //               </View>
  //             ))}
  //           </View>
  //         </View>
  //       </View>
  //     </HeaderScrollView>
  //     {!isMatched && (
  //       <View style={styles.reactionContainer}>
  //         <Reaction handleReact={handleReact} />
  //       </View>
  //     )}
  //   </> : userNotFound;
}

Person.defaultProps = {
  handleReact: () => {}
};

const styles = StyleSheet.create({
  headerStyle: {
    backgroundColor: "white"
  },
  imageStyle: {
    borderRadius: 10,
    width: "100%",
    height: 350,
    marginLeft: "auto",
    marginRight: "auto"
  },
  textStyle: {
    fontFamily: "DMSansRegular"
  },
  infoContainer: {
    margin: 20
  },
  basicInfoContainer: {
    marginTop: 20
  },
  basicInfoRow: {
    flexDirection: "row",
    display: "flex",
    alignItems: "center",
    marginBottom: 5
  },
  basicInfoRowIcon: {
    marginRight: 5,
    float: "left",
    color: "#6B5CBF"
  },
  basicInfoRowText: {
    fontFamily: "DMSansRegular",
    float: "right"
  },
  promptContainer: {
    fontFamily: "DMSansRegular",
    marginTop: 20
  },
  separator: {
    borderBottomWidth: 1,
    borderBottomColor: "#CED0CE", // Change the color as per your preference
    marginVertical: 29 // Adjust vertical spacing as needed
  },
  instagramContainer: {
    backgroundColor: "white",
    borderRadius: 8,
    paddingVertical: 10,
    paddingHorizontal: 10,
    width: "65%",
    marginVertical: 10,

    shadowColor: "#171717",
    shadowOffset: { width: 0, height: 0 },
    shadowOpacity: 0.2,
    shadowRadius: 4,

    flexDirection: "row",
    display: "flex",
    alignItems: "center"
  },
  mutualEventItemStyling: {
    marginRight: 10,
    marginTop: 10
  },
  reactionContainer: {
    position: "absolute",
    width: "100%",
    bottom: 70
  }
});
