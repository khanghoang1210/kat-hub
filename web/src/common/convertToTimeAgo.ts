export function timeAgo(date: string): string {
  const now = new Date();
  const past = new Date(date);
  const seconds = Math.floor((now.getTime() - past.getTime()) / 1000);

  let interval = Math.floor(seconds / 31536000);
  if (interval >= 1 && interval < 2) {
    return "1 year ago";
  }
  if (interval >= 2) {
    return `${interval} years ago`;
  }
  
  interval = Math.floor(seconds / 2592000);
  if (interval >= 1 && interval < 2) {
    return "1 month ago";
  }
  if (interval >= 2) {
    return `${interval} months ago`;
  }
  
  interval = Math.floor(seconds / 86400);
  if (interval >= 1 && interval < 2) {
    return "1 day ago";
  }
  if (interval >= 2) {
    return `${interval} days ago`;
  }
  
  interval = Math.floor(seconds / 3600);
  if (interval >= 1 && interval < 2) {
    return "1 hour ago";
  }
  if (interval >= 2) {
    return `${interval} hours ago`;
  }
  
  interval = Math.floor(seconds / 60);
  if (interval >= 1 && interval < 2) {
    return "1 minute ago";
  }
  if (interval >= 2) {
    return `${interval} minutes ago`;
  }
  
  return "just now";
}
