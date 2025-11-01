package utils

import (
	"fmt"
	"strings"

	"topinambur02.com/m/v2/model"
)

func FormatLesson(lesson model.Lesson) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("<b>%s</b>\n", lesson.Name))
	builder.WriteString(fmt.Sprintf("🕒 <i>%s</i>\n", lesson.TimeInterval))
	builder.WriteString(fmt.Sprintf("📅 %s\n", lesson.DateInterval))
	builder.WriteString(fmt.Sprintf("🏫 %s", lesson.Place))

	if len(lesson.Rooms) > 0 {
		builder.WriteString(fmt.Sprintf(" (ауд. %s)", strings.Join(lesson.Rooms, ", ")))
	}

	builder.WriteString(fmt.Sprintf("\n👨‍🏫 %s", strings.Join(lesson.Teachers, ", ")))

	if lesson.Link != nil && *lesson.Link != "" {
		builder.WriteString(fmt.Sprintf("\n🔗 <a href=\"%s\">Ссылка на занятие</a>", *lesson.Link))
	}

	return builder.String()
}
