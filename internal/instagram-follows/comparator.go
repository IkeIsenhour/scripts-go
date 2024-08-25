package insta

func CompareFollowersAgainstFollowing(i *InstagramRequest) ([]Users, error) {
	followerMap := make(map[string]bool)

	followers, err := i.GetAllFollowers()
	if err != nil {
		return nil, err
	}

	for _, v := range followers {
		followerMap[v.Username] = true
	}

	following, err := i.GetAllFollowing()
	if err != nil {
		return nil, err
	}

	var followingWhoDontFollowBack []Users
	for _, v := range following {
		if !followerMap[v.Username] {
			followingWhoDontFollowBack = append(followingWhoDontFollowBack, v)
		}
	}

	return followingWhoDontFollowBack, nil
}
