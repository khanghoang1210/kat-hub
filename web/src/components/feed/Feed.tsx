import React, { useEffect, useState } from "react";
import Leftbar from "./Leftbar";
import FriendItem from "./FriendItem";
import StatusBar from "./StatusBar";
import PostItem from "./PostItem";
import { Constants } from "@/common/constants";
import { Post } from "@/models/post";
import { timeAgo } from "@/common/convertToTimeAgo";
import { User } from "@/models/user";
import { PostService } from "@/services/PostService";
import UserService from "@/services/UserService";

type FeedProps = {};
const USER_ENDPOINT = "/users";
const Feed: React.FC<FeedProps> = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [currentUser, setCurrentUser] = useState<User>(Object);


  useEffect(() => {
    const postService = new PostService();
    const userService = new UserService();
    const fetchPostData = async ()=>{
      try {
        const posts = await postService.getPosts()
        setPosts(posts.data)
      }
      catch(error){
        throw error
      }
    }

    const fetchUserData = async () => {
      try {
        const user = await userService.getUserByID(12)
        setCurrentUser(user.data)
      } catch (error) {
        throw error;
      }
    }
    fetchPostData();
    fetchUserData();
  }, []);

  return (
    <div className="flex">
      <Leftbar userName={currentUser.userName} 
        title={currentUser.title}
        avatarUrl={currentUser.avatarUrl}
      />
      <div className="flex-grow">
        <StatusBar avatarUrl={currentUser.avatarUrl}/>
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
                userImage={post.author.avatarUrl}
              />
            );
          })}
        </div>
      </div>
    </div>
  );
};
export default Feed;
