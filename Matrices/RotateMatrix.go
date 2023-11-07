package matrices

//You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

//You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

//Traverse groups of four cells in the matrix, starting from the four corners
//swap the top left cell with the top right cell
//swap the top left cell with the bottom right cell
//swap the top left cell with the bottom left cell
//Move to the next group of four cells and repeat the above steps

