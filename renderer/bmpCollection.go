package renderer

type renderBMPCollection struct {
	BMPRects []*bmpRect
}

type Type_SpriteId uint32

func (collection *renderBMPCollection) AddBMP(bmp bmpRect) Type_SpriteId {
	for i := 0; i < len(collection.BMPRects); i++ {
		if collection.BMPRects[i] == nil {
			collection.BMPRects[i] = &bmp
			return Type_SpriteId(i)
		}
	}

	collection.BMPRects = append(collection.BMPRects, &bmp)
	return Type_SpriteId(len(collection.BMPRects) - 1)
}

func (collection *renderBMPCollection) Destroy() {
	for _, rect := range collection.BMPRects {
		if rect != nil {
			rect.DestroyBMPRect()
		}
	}

	collection.BMPRects = []*bmpRect{}
}

func (collection *renderBMPCollection) GetSprite(spriteId Type_SpriteId) *bmpRect {
	for index := range collection.BMPRects {
		if index == int(spriteId) {
			return collection.BMPRects[index]
		}
	}
	return nil
}
