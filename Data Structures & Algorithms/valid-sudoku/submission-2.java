class Solution {

	public boolean isValidSudoku(char[][] board) {
		//Each row must contain the digits 1-9 without duplicates.
		//Each column must contain the digits 1-9 without duplicates.
		//Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
		for(int i = 0; i < board.length; i++) {
			if(!rowChecker(board[i])) return false;
			int[] array = new int[board.length + 1];
			for(int j = 0; j < board[i].length; j++) {
				int current = Character.getNumericValue(board[j][i]);
				if(current < 1) continue;
				array[current]++;
				if(array[current] > 1) return false;
			}
		}
		if (!quadrantChecker(board)) return false;
		return true;

	}

	private static boolean rowChecker(char[] row) {
		int[] array = new int[row.length + 1];

		for(int i = 0; i < row.length; i++) {
			int current = Character.getNumericValue(row[i]);
			if(current < 1) continue;
			array[current]++;
			if(array[current] > 1) return false;
		}
		return true;
	}

	private static boolean quadrantChecker(char[][] board) {
		for (int row = 0; row < 9; row += 3) {
			for (int col = 0; col < 9; col += 3) {
				int[] seen = new int[10];

				for (int i = row; i < row + 3; i++) {
					for (int j = col; j < col + 3; j++) {
						int current = Character.getNumericValue(board[i][j]);
						if (current < 1) continue;

						seen[current]++;
						if (seen[current] > 1) return false;
					}
				}
			}
		}
		return true;
	}

}


