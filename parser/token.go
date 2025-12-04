package parser

type Token interface {
	GetText() string
}

type SQLToken struct {
	lexeme    string
	tokenType TokenType
}

type Pos struct {
	Line   int
	Column int
}

func NewSQLToken(text string, tokenType TokenType) *SQLToken {
	return &SQLToken{lexeme: text, tokenType: tokenType}
}

func (s *SQLToken) GetText() string {
	return s.lexeme
}

var operators = map[string]TokenType{
	";":   SCOL,
	".":   DOT,
	"(":   OPEN_PAR,
	")":   CLOSE_PAR,
	",":   COMMA,
	"=":   ASSIGN,
	"*":   STAR,
	"+":   PLUS,
	"-":   MINUS,
	"~":   TILDE,
	"||":  PIPE2,
	"/":   DIV,
	"%":   MOD,
	"<<":  LT2,
	">>":  GT2,
	"&":   AMP,
	"|":   PIPE,
	"<":   LT,
	"<=":  LT_EQ,
	">":   GT,
	">=":  GT_EQ,
	"==":  EQ,
	"!=":  NOT_EQ1,
	"<>":  NOT_EQ2,
	"->":  JPTR,
	"->>": JPTR2,
}

type TokenType int

const (
	TokenUnknown TokenType = iota

	// punctuation / operators
	SCOL
	DOT
	OPEN_PAR
	CLOSE_PAR
	COMMA
	ASSIGN
	STAR
	PLUS
	MINUS
	TILDE
	PIPE2
	DIV
	MOD
	LT2
	GT2
	AMP
	PIPE
	LT
	LT_EQ
	GT
	GT_EQ
	EQ
	NOT_EQ1
	NOT_EQ2
	JPTR
	JPTR2

	// keywords (kept with trailing underscore as in grammar)
	ABORT_
	ACTION_
	ADD_
	AFTER_
	ALL_
	ALTER_
	ANALYZE_
	AND_
	AS_
	ASC_
	ATTACH_
	AUTOINCREMENT_
	BEFORE_
	BEGIN_
	BETWEEN_
	BY_
	CASCADE_
	CASE_
	CAST_
	CHECK_
	COLLATE_
	COLUMN_
	COMMIT_
	CONFLICT_
	CONSTRAINT_
	CREATE_
	CROSS_
	CURRENT_DATE_
	CURRENT_TIME_
	CURRENT_TIMESTAMP_
	DATABASE_
	DEFAULT_
	DEFERRABLE_
	DEFERRED_
	DELETE_
	DESC_
	DETACH_
	DISTINCT_
	DROP_
	EACH_
	ELSE_
	END_
	ESCAPE_
	EXCEPT_
	EXCLUSIVE_
	EXISTS_
	EXPLAIN_
	FAIL_
	FOR_
	FOREIGN_
	FROM_
	FULL_
	GLOB_
	GROUP_
	HAVING_
	IF_
	IGNORE_
	IMMEDIATE_
	IN_
	INDEX_
	INDEXED_
	INITIALLY_
	INNER_
	INSERT_
	INSTEAD_
	INTERSECT_
	INTO_
	IS_
	ISNULL_
	JOIN_
	KEY_
	LEFT_
	LIKE_
	LIMIT_
	MATCH_
	MATERIALIZED_
	NATURAL_
	NO_
	NOT_
	NOTNULL_
	NULL_
	OF_
	OFFSET_
	ON_
	OR_
	ORDER_
	OUTER_
	PLAN_
	PRAGMA_
	PRIMARY_
	QUERY_
	RAISE_
	RECURSIVE_
	REFERENCES_
	REGEXP_
	REINDEX_
	RELEASE_
	RENAME_
	REPLACE_
	RESTRICT_
	RETURNING_
	RIGHT_
	ROLLBACK_
	ROW_
	ROWS_
	ROWID_
	SAVEPOINT_
	SELECT_
	SET_
	STRICT_
	TABLE_
	TEMP_
	TEMPORARY_
	THEN_
	TO_
	TRANSACTION_
	TRIGGER_
	UNION_
	UNIQUE_
	UPDATE_
	USING_
	VACUUM_
	VALUES_
	VIEW_
	VIRTUAL_
	WHEN_
	WHERE_
	WITH_
	WITHOUT_
	OVER_
	PARTITION_
	RANGE_
	PRECEDING_
	UNBOUNDED_
	CURRENT_
	FOLLOWING_
	RANK_
	GENERATED_
	ALWAYS_
	STORED_
	TRUE_
	FALSE_
	WINDOW_
	NULLS_
	FIRST_
	LAST_
	FILTER_
	GROUPS_
	EXCLUDE_
	TIES_
	OTHERS_
	DO_
	NOTHING_

	// literals / identifiers / parameters
	IDENTIFIER
	NUMERIC_LITERAL
	BIND_PARAMETER
	STRING_LITERAL
	BLOB_LITERAL

	// comments / whitespace (often emitted on hidden channel)
	SINGLE_LINE_COMMENT
	MULTILINE_COMMENT
	SPACES

	// fallback
	UNEXPECTED_CHAR
)
