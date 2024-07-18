import React, { useEffect, useState } from "react";
import Leftbar from "./Leftbar";
import FriendItem from "./FriendItem";
import StatusBar from "./StatusBar";
import PostItem from "./PostItem";
import { Constants } from "@/common/constants";
import { Post } from "@/models/post";
import { timeAgo } from "@/common/convertToTimeAgo";
import { User } from "@/models/user";

type FeedProps = {};
const POST_ENDPOINT = "/posts";
const USER_ENDPOINT = "/users";
const Feed: React.FC<FeedProps> = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [currentUser, setCurrentUser] = useState<User>();
  useEffect(() => {
    async function fetchData() {
      try {
        const res = await fetch(Constants.API_URL + POST_ENDPOINT, {
          method: "GET",
          mode: "cors",
          headers: {
            Accept: "application/json, text/plain",
            "Content-Type": "application/json;charset=UTF-8",
          },
        });

        const data = await res.json();
        setPosts(data.data);
      } catch (error) {
        throw error;
      }
    }

    async function fetchUserData() {
      try {
        const res = await fetch(Constants.API_URL + USER_ENDPOINT + "/12", {
          method: "GET",
          mode: "cors",
          headers: {
            Accept: "application/json, text/plain",
            "Content-Type": "application/json;charset=UTF-8",
          },
        });

        const data = await res.json();
        console.log(data.data);
        setCurrentUser(data.data);
      } catch (error) {
        throw error;
      }
    }
    fetchData();
    fetchUserData();
  }, []);

  return (
    <div className="flex">
      <Leftbar userName={currentUser?.userName} 
        title={currentUser?.title}
        avatarUrl={currentUser?.avatarUrl}
      />
      <div className="flex-grow">
        <StatusBar />
        <FriendItem />
        <div>
          {posts.map((post, idx) => {
            return (
              <PostItem
                key={idx}
                username={post.author.userName}
                userTitle={post.author.title}
                postTime={timeAgo(post.createdAt.toString())}
                postContent={post.textContent}
                userImage="/assets/icons/user.svg"
              />
            );
          })}
        </div>
      </div>
    </div>
  );
};
export default Feed;
