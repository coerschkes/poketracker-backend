package external

const (
	selectUserQuery = "SELECT id FROM useraccount WHERE firebaseuid = $1"
)

type UserRepository interface {
	Find(firebaseUid string) (int, error)
}

type UserRepositoryImpl struct {
	connector *DatabaseConnector
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{connector: NewDatabaseConnector()}
}

func (i *UserRepositoryImpl) Find(firebaseUid string) (int, error) {
	query, err := i.connector.Query(selectUserQuery, NewUserMapper(), firebaseUid)
	if err != nil {
		return 0, err
	}
	return query.(int), nil
}
