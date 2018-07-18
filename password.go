package validate

type constraint struct {
	val	string
}

var (
	LowerCase 		= constraint{"LOWER_CASE"}
	UpperCase 		= constraint{"UPPER_CASE"}
	Number	  		= constraint{"NUMBER"}
	SpecialCharacter	= constraint{"SPECIAL_CHARACTER"}
)

type PasswordConfigs struct {
	MustContain 	[]constraint
	MinLength	int
	MaxLength	int	
}

type PasswordConfig func(*PasswordConfigs)

func MustContain(constraints []constraint) PasswordConfig {
	return func(args *PasswordConfigs) {
		args.MustContain = constraints  
	}
}
