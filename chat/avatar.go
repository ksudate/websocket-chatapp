package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL はAvatarインスタンスがアバターのURLを返すことができない場合のエラー
var ErrNoAvatarURL = errors.New("chat: アバターのURLが取得できません")

// Avatar はユーザーのプロフィール画像を表す型
type Avatar interface {
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars expects .
type TryAvatars []Avatar

// GetAvatarURL expects .
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar expects .
type AuthAvatar struct{}

// UseAuthAvatar expects .
var UseAuthAvatar AuthAvatar

// GetAvatarURL expects .
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar expects .
type GravatarAvatar struct{}

// UseGravatar expects .
var UseGravatar GravatarAvatar

// GetAvatarURL expects .
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar expects .
type FileSystemAvatar struct{}

// UseFileSystemAvatar expects .
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL expects .
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
