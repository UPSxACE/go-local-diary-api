import { fromEvent } from "rxjs";

console.log("This script will be in every page!");

fromEvent(document, "click").subscribe(() => {
  console.log("A click happened! The rxjs library is working!");
});
